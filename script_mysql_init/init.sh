#!/bin/sh
# wx server
echo "mysql check"
mysql -h mariadb -uroot -p123123 -e "CREATE USER 'wx'@'%' IDENTIFIED BY 'rWk1hvqMT62K2JYH';"
mysql -h mariadb -uroot -p123123 -e "CREATE DATABASE IF NOT EXISTS wx;"
mysql -h mariadb -uroot -p123123 -e "GRANT ALL ON wx.* TO 'wx'@'%';"


