#!/usr/bin/env bash
# Explicit Spanish development sample data.
set -euo pipefail
PROJECT="$(cd "$(dirname "$0")/.." && pwd)"
exec "$PROJECT/scripts/dev-seed-runner.sh" es
