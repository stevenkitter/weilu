#!/bin/sh

# save the path
workspace=$PWD

# set env
export WX_MYSQL_USER=wx
export WX_MYSQL_PASSWORD=rWk1hvqMT62K2JYH
export WX_DATABASE=wx
export WX_SERVER_ADDRESS=localhost:51001
## docker
export DOCKER_PASSWORD=wangxu226

# # build image and push to tencent 
# # ----------------------------------------------------------------
# cd $PWD/services/wx
# CGO_ENABLED=0 GOOS=linux go build -o app .
# docker build -t ccr.ccs.tencentyun.com/julu/weilu_wx .
# docker login -u 100000776220 -p $DOCKER_PASSWORD ccr.ccs.tencentyun.com
# docker push ccr.ccs.tencentyun.com/julu/weilu_wx

# # run the docker-compose
# # ----------------------------------------------------------------
# cd $workspace
# docker network list -f name=internal | wc -l | awk '$0==1{cmd="docker network create internal";system(cmd)}'
# docker-compose -f docker-compose.yml up -d


