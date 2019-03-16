docker build -t ansible .
docker rmi $(docker images -f "dangling=true" -q)
docker-compose up -d
docker exec -it ansible ash
