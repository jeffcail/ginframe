#!/bin/bash

echo " ========= 开始上传 server-user ========= "

function scp_file1 {
    local file=$1
    local passwd="112233"
    expect -c"
        spawn scp ${file} root@192.168.0.72:/root/server-user
        expect {
          \"*password\" {set timeout 300; send \"${passwd}\r\";}
        }
        expect eof"
}


function scp_file2 {
    local file=$1
    local passwd="112233"
    expect -c"
        spawn scp ${file} root@192.168.0.73:/root/server-user
        expect {
          \"*password\" {set timeout 300; send \"${passwd}\r\";}
        }
        expect eof"
}

scp_file1 "server-user"

echo " ========= server-user1 服务上传完成 ========= "

#sleep 5

echo " ========= 开始上传 server-user2 ========= "

scp_file2 "server-user"

echo " ========= server-user2 服务上传完成 ========= "

