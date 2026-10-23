#!/usr/bin/env bash
# Explicit English development sample data.
set -euo pipefail
PROJECT="$(cd "$(dirname "$0")/.." && pwd)"
exec "$PROJECT/scripts/dev-seed-runner.sh" en
