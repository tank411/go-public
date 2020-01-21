#!/bin/bash
echo "start build for linux/windows/mac!"

echo "build server"
#cd src/server
# 添加图标.教程: https://blog.csdn.net/u014633966/article/details/82984037
rsrc -ico icon.ico -o main.syso

# 编译各版本服务器端
echo "build server to linux"
export  GOOS=linux
export  GOARCH=amd64
go build -ldflags "-s -w" -o ./linux/14_day
echo "build server to windows"
export  GOOS=windows
export  GOARCH=amd64
go build -ldflags "-s -w -H windowsgui" -o ./windows/14_day.exe
echo "build server to mac"
export  GOOS=darwin
export  GOARCH=amd64

go build -ldflags "-s -w"  -o ./mac/14_day
echo "build server finished!"

echo "build client"
cd ../client
# 添加图标.教程: https://blog.csdn.net/u014633966/article/details/82984037
rsrc -ico icon.ico -o main.syso


# 加壳
echo "start upx ..."
cd ../_pkg/bin
upx -9 -k windows/xdataops.exe
upx -9 -k linux/xdataops
upx -9 -k mac/xdataops

echo "app build to path[ go-web/bin ]"

