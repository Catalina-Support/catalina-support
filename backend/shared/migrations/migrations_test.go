package migrations

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverOnlyStrictProductionMigrationsInSemverOrder(t *testing.T) {
	dir := t.TempDir()
	for name, body := range map[string]string{
		"v1.0.10.sql": "SELECT 10", "v1.0.2.sql": "SELECT 2",
		"v1.0.0_dev_en.sql": "DELETE FROM users", "v1.0.0_dev_es.sql": "DELETE FROM users",
		"notes.sql": "SELECT 1",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	files, err := discover(dir, semver{1, 0, 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 || files[0].name != "v1.0.2.sql" || files[1].name != "v1.0.10.sql" {
		t.Fatalf("unexpected files: %#v", files)
	}
}

func TestDiscoverRejectsMigrationNewerThanApplication(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "v1.1.0.sql"), []byte("SELECT 1"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := discover(dir, semver{1, 0, 0}); err == nil {
		t.Fatal("expected future migration error")
	}
}

func TestDiscoverRejectsInvalidVersionedFilename(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "v1.01.0.sql"), []byte("SELECT 1"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := discover(dir, semver{1, 1, 0}); err == nil {
		t.Fatal("expected invalid filename error")
	}
}

func TestValidateHistoryAllowsOnlyCurrentDevelopmentChecksumChange(t *testing.T) {
	files := []migration{{version: semver{1, 0, 0}, name: "v1.0.0.sql", checksum: "new"}}
	if err := validateHistory(map[string]string{"1.0.0": "old"}, files, semver{1, 0, 0}, true); err != nil {
		t.Fatal(err)
	}
	if err := validateHistory(map[string]string{"1.0.0": "old"}, files, semver{1, 0, 0}, false); err == nil {
		t.Fatal("published checksum change must fail")
	}
}

func TestParseVersionIsStrict(t *testing.T) {
	for _, invalid := range []string{"1.0", "v1.0.0", "1.01.0", "1.0.0-beta"} {
		if _, err := parseVersion(invalid); err == nil {
			t.Fatalf("accepted %q", invalid)
		}
	}
}
