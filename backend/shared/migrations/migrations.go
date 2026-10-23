// Package migrations applies the versioned SQL schema before the HTTP server starts.
package migrations

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"catalina-support/backend/shared/version"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const advisoryLockID int64 = 0x434154414C494E41

var migrationName = regexp.MustCompile(`^v([0-9]+)\.([0-9]+)\.([0-9]+)\.sql$`)

type Config struct {
	Host, Port, User, Password, Database string
	Directory, BackupDirectory           string
	Version                              string
	Development                          bool
	Now                                  func() time.Time
}

type migration struct {
	version  semver
	name     string
	path     string
	content  []byte
	checksum string
}

type semver struct{ major, minor, patch int }

func (v semver) String() string { return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch) }
func (v semver) compare(other semver) int {
	if v.major != other.major {
		return cmp(v.major, other.major)
	}
	if v.minor != other.minor {
		return cmp(v.minor, other.minor)
	}
	return cmp(v.patch, other.patch)
}
func cmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func ConfigFromEnv() Config {
	return Config{
		Host: os.Getenv("POSTGRES_HOST"), Port: os.Getenv("PGPORT"), User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"), Database: os.Getenv("POSTGRES_DB"),
		Directory:       envOr("MIGRATIONS_DIR", "/migrations"),
		BackupDirectory: envOr("MIGRATION_BACKUP_DIR", "/backups"),
		Version:         version.Version, Development: version.Development, Now: time.Now,
	}
}

func Run(ctx context.Context, cfg Config) error {
	if cfg.Now == nil {
		cfg.Now = time.Now
	}
	target, err := parseVersion(cfg.Version)
	if err != nil {
		return fmt.Errorf("application version: %w", err)
	}
	files, err := discover(cfg.Directory, target)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("no versioned migration files found")
	}

	db, err := sql.Open("pgx", dsn(cfg))
	if err != nil {
		return err
	}
	defer db.Close()
	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("connect to PostgreSQL: %w", err)
	}
	defer conn.Close()
	if _, err = conn.ExecContext(ctx, "SELECT pg_advisory_lock($1)", advisoryLockID); err != nil {
		return fmt.Errorf("acquire migration lock: %w", err)
	}
	defer conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock($1)", advisoryLockID)

	applied, tableExists, err := readApplied(ctx, conn)
	if err != nil {
		return err
	}
	if err := validateHistory(applied, files, target, cfg.Development); err != nil {
		return err
	}
	pending := pendingFiles(applied, files, cfg.Development, target)
	if len(pending) == 0 {
		fmt.Printf("schema already at v%s; no backup or migration needed\n", target)
		return nil
	}
	from := "none"
	for v := range applied {
		if parsed, e := parseVersion(v); e == nil && (from == "none" || parsed.compare(mustVersion(from)) > 0) {
			from = v
		}
	}
	backup, err := createBackup(ctx, cfg, from, target.String())
	if err != nil {
		return err
	}
	fmt.Printf("backup verified: %s\n", backup)

	for _, file := range pending {
		if err := apply(ctx, conn, file, tableExists); err != nil {
			return fmt.Errorf("apply %s: %w", file.name, err)
		}
		tableExists = true
		fmt.Printf("migration applied: %s (%s)\n", file.name, file.checksum)
	}
	pruneBackups(cfg.BackupDirectory, cfg.Now().Add(-14*24*time.Hour))
	return nil
}

func discover(directory string, target semver) ([]migration, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("read migrations: %w", err)
	}
	var result []migration
	seen := map[string]bool{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		match := migrationName.FindStringSubmatch(entry.Name())
		if match == nil && strings.HasPrefix(entry.Name(), "v") && strings.HasSuffix(entry.Name(), ".sql") && !strings.HasSuffix(entry.Name(), "_dev_en.sql") && !strings.HasSuffix(entry.Name(), "_dev_es.sql") {
			return nil, fmt.Errorf("invalid migration filename %s", entry.Name())
		}
		if match == nil {
			continue
		}
		v, err := parseVersion(strings.TrimSuffix(strings.TrimPrefix(entry.Name(), "v"), ".sql"))
		if err != nil {
			return nil, fmt.Errorf("invalid migration filename %s: %w", entry.Name(), err)
		}
		if v.compare(target) > 0 {
			return nil, fmt.Errorf("migration %s is newer than application v%s", entry.Name(), target)
		}
		if seen[v.String()] {
			return nil, fmt.Errorf("duplicate migration version %s", v)
		}
		seen[v.String()] = true
		body, err := os.ReadFile(filepath.Join(directory, entry.Name()))
		if err != nil {
			return nil, err
		}
		sum := sha256.Sum256(body)
		result = append(result, migration{v, entry.Name(), filepath.Join(directory, entry.Name()), body, hex.EncodeToString(sum[:])})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].version.compare(result[j].version) < 0 })
	return result, nil
}

func readApplied(ctx context.Context, conn *sql.Conn) (map[string]string, bool, error) {
	var exists bool
	if err := conn.QueryRowContext(ctx, "SELECT to_regclass('public.schema_migrations') IS NOT NULL").Scan(&exists); err != nil {
		return nil, false, err
	}
	result := map[string]string{}
	if !exists {
		return result, false, nil
	}
	rows, err := conn.QueryContext(ctx, "SELECT version, checksum FROM schema_migrations")
	if err != nil {
		return nil, true, err
	}
	defer rows.Close()
	for rows.Next() {
		var v, c string
		if err := rows.Scan(&v, &c); err != nil {
			return nil, true, err
		}
		result[v] = c
	}
	return result, true, rows.Err()
}

func validateHistory(applied map[string]string, files []migration, target semver, development bool) error {
	byVersion := map[string]migration{}
	for _, file := range files {
		byVersion[file.version.String()] = file
	}
	for v, checksum := range applied {
		parsed, err := parseVersion(v)
		if err != nil {
			return fmt.Errorf("invalid recorded migration version %q", v)
		}
		if parsed.compare(target) > 0 {
			return fmt.Errorf("database v%s is newer than application v%s", parsed, target)
		}
		file, ok := byVersion[v]
		if !ok {
			return fmt.Errorf("recorded migration v%s is missing from this release", v)
		}
		if checksum != file.checksum && !(development && parsed.compare(target) == 0) {
			return fmt.Errorf("checksum mismatch for immutable migration %s", file.name)
		}
	}
	return nil
}

func pendingFiles(applied map[string]string, files []migration, development bool, target semver) []migration {
	var result []migration
	for _, file := range files {
		checksum, ok := applied[file.version.String()]
		if !ok || (development && file.version.compare(target) == 0 && checksum != file.checksum) {
			result = append(result, file)
		}
	}
	return result
}

func apply(ctx context.Context, conn *sql.Conn, file migration, tableExists bool) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if !tableExists {
		if _, err = tx.ExecContext(ctx, `CREATE TABLE schema_migrations (version text PRIMARY KEY, filename text NOT NULL UNIQUE, checksum char(64) NOT NULL, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
			return err
		}
	}
	if _, err = tx.ExecContext(ctx, string(file.content)); err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO schema_migrations(version, filename, checksum, applied_at) VALUES ($1,$2,$3,now()) ON CONFLICT (version) DO UPDATE SET filename=EXCLUDED.filename, checksum=EXCLUDED.checksum, applied_at=EXCLUDED.applied_at`, file.version.String(), file.name, file.checksum)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func createBackup(ctx context.Context, cfg Config, from, to string) (string, error) {
	if err := os.MkdirAll(cfg.BackupDirectory, 0o750); err != nil {
		return "", fmt.Errorf("create backup directory: %w", err)
	}
	name := fmt.Sprintf("%s_v%s_to_v%s_%s.dump", cfg.Database, from, to, cfg.Now().UTC().Format("20060102_150405"))
	path := filepath.Join(cfg.BackupDirectory, name)
	args := []string{"-h", cfg.Host, "-p", cfg.Port, "-U", cfg.User, "-d", cfg.Database, "-Fc", "-f", path}
	cmd := exec.CommandContext(ctx, "pg_dump", args...)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+cfg.Password)
	if output, err := cmd.CombinedOutput(); err != nil {
		os.Remove(path)
		return "", fmt.Errorf("pg_dump: %w: %s", err, strings.TrimSpace(string(output)))
	}
	if info, err := os.Stat(path); err != nil || info.Size() == 0 {
		os.Remove(path)
		return "", errors.New("pg_dump produced an empty backup")
	}
	verify := exec.CommandContext(ctx, "pg_restore", "--list", path)
	if output, err := verify.CombinedOutput(); err != nil {
		return "", fmt.Errorf("verify backup: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return path, nil
}

func pruneBackups(directory string, before time.Time) {
	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		info, err := entry.Info()
		if err == nil && !entry.IsDir() && strings.HasSuffix(entry.Name(), ".dump") && info.ModTime().Before(before) {
			_ = os.Remove(filepath.Join(directory, entry.Name()))
		}
	}
}

func parseVersion(raw string) (semver, error) {
	parts := strings.Split(raw, ".")
	if len(parts) != 3 {
		return semver{}, fmt.Errorf("%q is not strict SemVer", raw)
	}
	values := [3]int{}
	for i, part := range parts {
		if part == "" || (len(part) > 1 && part[0] == '0') {
			return semver{}, fmt.Errorf("%q is not strict SemVer", raw)
		}
		n, err := strconv.Atoi(part)
		if err != nil || n < 0 {
			return semver{}, fmt.Errorf("%q is not strict SemVer", raw)
		}
		values[i] = n
	}
	return semver{values[0], values[1], values[2]}, nil
}
func mustVersion(raw string) semver { v, _ := parseVersion(raw); return v }
func dsn(cfg Config) string {
	address := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host + ":" + cfg.Port,
		Path:     cfg.Database,
		RawQuery: "sslmode=disable",
	}
	return address.String()
}
func envOr(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}
