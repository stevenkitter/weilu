#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$POSTGRES_PASSWORD mysql -h"$host" -uroot -p123123; do
  >&2 echo "Mysql is unavailable - sleeping"
  sleep 1
done

>&2 echo "mysql is up - executing command"
exec $cmd