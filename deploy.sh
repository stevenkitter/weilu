#!/bin/sh
docker login -u $DOCKER_TENCENT_USER -p $DOCKER_TENCENT_PASSWORD ccr.ccs.tencentyun.com
docker pull ccr.ccs.tencentyun.com/julu/weilu_wx
docker-compose down --remove-orphans && docker-compose up -d