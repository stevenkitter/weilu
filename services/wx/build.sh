
CGO_ENABLED=0 GOOS=linux go build -o app .
docker build -t ccr.ccs.tencentyun.com/julu/weilu_wx .
docker login -u 100000776220 -p wangxu226 ccr.ccs.tencentyun.com
docker push ccr.ccs.tencentyun.com/julu/weilu_wx



