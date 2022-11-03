#!/bin/sh

set -e # run immediately

echo "run db migration"
echo "$DB_URL"
/app/migrate -path /app/migrations -database "$DB_URL" -verbose up

echo "start the app"
exec "$@" # takes all the parameters passed to the script and run it