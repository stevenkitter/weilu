#!/bin/sh
# wx server
echo "mysql check"
sudo service mysql start
mysql -h 127.0.0.1 -uroot -p123123 -e "CREATE USER 'wx'@'%' IDENTIFIED BY 'rWk1hvqMT62K2JYH';"
mysql -h 127.0.0.1 -uroot -p123123 -e "CREATE DATABASE IF NOT EXISTS wx;"
mysql -h 127.0.0.1 -uroot -p123123 -e "GRANT ALL ON wx.* TO 'wx'@'%';"


