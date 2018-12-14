#!/bin/sh
# wx server
echo "mysql check"
mysql -uroot -p123123 -e "CREATE USER 'wx'@'%' IDENTIFIED BY 'rWk1hvqMT62K2JYH';"
mysql -uroot -p123123 -e "CREATE DATABASE IF NOT EXISTS wx;"
mysql -uroot -p123123 -e "GRANT ALL ON wx.* TO 'wx'@'%';"


