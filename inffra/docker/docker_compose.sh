date > date.txt && docker-compose up -d
docker run -p 9000:9000 --link docker_mysql_1:mysql  -it docker_common/api bash
