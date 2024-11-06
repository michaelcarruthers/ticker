#!/usr/bin/env sh
#
# Perform setup of the devcontainer
#
# Setup installs the necessary dependencies for developing ticker.
# This includes a mock API server that serves `response.json

set -euo pipefail

MOCK_HTTP_HOST="${TICKER_HTTP_HOST:-"localhost"}"
MOCK_HTTP_PORT="${TICKER_HTTP_PORT:-"9090"}"
MOCK_HTTP_ROOT="${TICKER_HTTP_ROOT:-"/tmp/ticker"}"
MOCK_HTTP_UTIL="${TICKER_HTTP_START:-"/tmp/start-mock-api.sh"}"

apk --no-cache --quiet add caddy curl git

cat <<EOF>> /tmp/start-mock-api.sh
caddy file-server \
  --browse \
  --listen "${MOCK_HTTP_HOST}:${MOCK_HTTP_PORT}" \
  --root "${MOCK_HTTP_ROOT}" \
  &> /dev/null &
EOF
chmod 755 "${MOCK_HTTP_UTIL}"

echo "Start Mock API server with: ${MOCK_HTTP_UTIL}"