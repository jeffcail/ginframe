#! /bin/bash

# nacos服务地址
nacos=192.168.0.40


if [ ! -f 'server-user' ]; then
  echo 文件不存在! 待添加的安装包: 'server-user'
  exit
fi

echo "server-user..."
sleep 3
docker stop server-user

sleep 2
docker rm server-user

docker rmi server-user
echo ""

echo "server-user packing..."
sleep 3
docker build -t server-user .
echo ""

echo "server-user running..."
sleep 3

docker run \
  -p 9091:9091 \
  -p 9092:9092 \
  -p 9093:9093 \
  -p 9094:9094 \
  --name server-user \
  --net host \
  -v /mnt/server-user:/root/server-user/log \
  -v /etc/localtime:/etc/localtime \
  -v /data/leveldb_data:/root/server-user/leveldb_data \
  -d server-user \
  server-user -ip $nacos -p 7848 -c server-user.yml


docker logs -f server-user | sed '/Started server-user application/q'

echo ""