#!/bin/sh
docker login -u 100000776220 -p wangxu226 ccr.ccs.tencentyun.com
docker pull ccr.ccs.tencentyun.com/julu/weilu_wx
docker-compose down && docker-compose up -d