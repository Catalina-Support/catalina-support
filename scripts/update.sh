#!/usr/bin/env bash
set -euo pipefail

REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
LOG_DIR="${CATALINA_UPDATE_LOG_DIR:-/srv/catalina-support/_logs}"
mkdir -p "$LOG_DIR"
LOG_FILE="$LOG_DIR/update-$(date -u +%Y%m%d_%H%M%S).log"
exec > >(tee -a "$LOG_FILE") 2>&1

finish() {
  status=$?
  if [ "$status" -eq 0 ]; then result=success; else result=failure; fi
  echo "update_finished=$(date -u -Is) result=$result status=$status"
}
trap finish EXIT

echo "update_started=$(date -u -Is)"
if [ "$(git -C "$REPO_DIR" branch --show-current)" != "main" ]; then
  echo "ERROR: la actualización sólo se ejecuta desde la rama main." >&2
  exit 1
fi
if [ -n "$(git -C "$REPO_DIR" status --porcelain)" ]; then
  echo "ERROR: el árbol tiene cambios locales; no se modificó nada." >&2
  exit 1
fi

OLD_COMMIT="$(git -C "$REPO_DIR" rev-parse HEAD)"
git -C "$REPO_DIR" fetch origin main
if ! git -C "$REPO_DIR" merge-base --is-ancestor HEAD origin/main; then
  echo "ERROR: main y origin/main divergen; no se puede avanzar de forma rápida." >&2
  exit 1
fi
git -C "$REPO_DIR" merge --ff-only origin/main
NEW_COMMIT="$(git -C "$REPO_DIR" rev-parse HEAD)"
echo "source_commit=$OLD_COMMIT target_commit=$NEW_COMMIT"

if ! "$REPO_DIR/scripts/prod-build.sh"; then
  echo "ERROR: la construcción o actualización falló; revise $LOG_FILE." >&2
  exit 1
fi

echo "log=$LOG_FILE"
