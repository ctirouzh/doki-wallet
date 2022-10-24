#!/usr/bin/env bash

set -e

CMD=$(which app)
PORT="${PORT:-50051}"
DB_HOST="${DB_HOST:-db}"
DB_PORT="${DB_PORT:-3306}"
DB_NAME="${DB_NAME:-doki_wallet}"
DB_USER="${DB_USER:-doki}"
DB_PASS="${DB_PASS}"

if [[ "$1" == "" ]]; then
    exec "${CMD}" -port="${PORT}"      \
                  -dbhost="${DB_HOST}" \
                  -dbport="${DB_PORT}" \
                  -dbname="${DB_NAME}" \
                  -dbuser="${DB_USER}" \
                  -dbpass="${DB_PASS}"
else
    exec "$@"
fi