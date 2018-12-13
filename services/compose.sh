docker network rm internal && docker network create internal
docker-compose -f services/docker-compose-test.yml up -d