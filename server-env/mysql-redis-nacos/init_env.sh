#! /bin/bash

# chmod +x ./deploy_docker.sh

# echo "=====>>>>> 开始安装 docker..." && sleep 1 && echo ""
# ./deploy_docker.sh

# echo "=====>>>>> docker安装完毕 " && sleep 1

# rm -rf ./deploy_docker.sh
# rm -rf ./daemon.json

echo "=====>>>>> 准备安装并启动 mysql " && sleep 1 && echo ""
docker-compose -f mysql.yml up -d

sleep 3 && docker ps && echo ""

echo "=====>>>>> mysql启动成功 "
echo "=====>>>>> 浏览器访问7880端口，数据库密码请查询 mysql.yml " && sleep 2

echo "=====>>>>> 准备安装并启动 rds " && sleep 1 && echo ""

docker-compose -f redis.yml up -d

echo "=====>>>>> redis启动成功" && sleep 1 && echo ""


echo "=====>>>>> 准备安装并启动 nacos " && sleep 1 && echo ""

docker-compose -f nacos-server.yml up -d

echo "=====>>>>> nacos启动成功" && sleep 1 && echo ""
