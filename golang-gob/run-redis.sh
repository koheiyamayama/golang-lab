#!/bin/bash

# コンテナ名を定義します
container_name="my_redis_container"

# コンテナがすでに存在する場合は削除します
if [ $(docker ps -a -f name=$container_name | grep -w $container_name | wc -l) -eq 1 ]; then
  echo "Container $container_name exists, removing..."
  docker rm -f $container_name
fi

# Redisイメージを起動します
echo "Starting Redis container..."
docker run --name $container_name -d -p 6379:6379 redis

echo "Redis container started with name $container_name"

# Redisコンテナでmonitorコマンドを実行します
echo "Running monitor command in Redis container..."
docker exec -it $container_name redis-cli monitor
