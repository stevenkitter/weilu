#!/bin/sh
# wait.sh

set -e

until mysql -h mariadb -uroot -p123123 -e "SELECT 1"; do
  >&2 echo "Mysql is unavailable - sleeping"
  sleep 1
done

>&2 echo "Mysql is up - executing command"
