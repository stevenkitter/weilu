#!/bin/sh
# wait.sh

set -e

until mysql -h 127.0.0.1 -uroot -p123123; do
  >&2 echo "Mysql is unavailable - sleeping"
  sleep 1
done

>&2 echo "Mysql is up - executing command"
