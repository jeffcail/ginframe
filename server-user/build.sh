#!/usr/bin/env bash
echo "开始编译"
buildFileName="server-user"
BuildTime=`date +'%Y.%m.%d.%H:%M:%S'`
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ${buildFileName}
echo "编译完成, 编译时间${BuildTime}"
echo "开始压缩文件"
upx -9 ${buildFileName}
echo "完成压缩"
