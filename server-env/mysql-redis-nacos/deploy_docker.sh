#! /bin/bash

yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine

yum install -y yum-utils device-mapper-persistent-data lvm2

yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

yum install docker-ce docker-ce-cli containerd.io -y

echo "=====>>>>> 启动 docker..." && sleep 1
systemctl start docker

echo "=====>>>>> 启动成功..." && sleep 1

echo "=====>>>>> test docker..." && sleep 1
docker pull hello-world

docker rmi hello-world

echo "=====>>>>> install docker-compose..." && sleep 1
curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
# curl -L "https://get.daocloud.io/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose 国内镜像源

chmod +x /usr/local/bin/docker-compose

ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

echo "=====>>>>> test docker-compose..." && sleep 1
docker-compose --version

echo "=====>>>>> 安装成功" && sleep 1

exit
