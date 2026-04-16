#!/bin/bash

# 定义国内镜像前缀
MIRROR_PREFIX="docker.m.daocloud.io/"

./stop.sh

# 1. 启动 MySQL（添加前缀）
echo "正在拉取并启动 MySQL..."
docker run --name fabrictrace-mysql -p 3337:3306 \
  -e MYSQL_ROOT_PASSWORD=fabrictrace -d \
  ${MIRROR_PREFIX}mysql:8

# 2. 检查并拉取 Fabric 镜像
image_versions=("2.5")
images=("hyperledger/fabric-tools" "hyperledger/fabric-peer" "hyperledger/fabric-orderer" "hyperledger/fabric-ccenv" "hyperledger/fabric-baseos")

for image in "${images[@]}"
do
    for version in "${image_versions[@]}"
    do
        FULL_IMAGE_NAME="${MIRROR_PREFIX}${image}:${version}"
        LOCAL_NAME="${image}:${version}"

        if ! docker images -a | grep "$image" | grep "$version" &> /dev/null
        then
            echo "镜像 $LOCAL_NAME 不存在，从私有源拉取..."
            docker pull "$FULL_IMAGE_NAME"
            # 关键步骤：拉取后打上 tag，这样 network.sh 脚本就能通过原名直接找到镜像
            docker tag "$FULL_IMAGE_NAME" "$LOCAL_NAME"
        fi
    done
done

# 启动区块链网络、创建通道
./network.sh up createChannel
# 部署链码，使用trace链码
./network.sh deployCC -ccn trace -ccp ../chaincode -ccl go
cp -r organizations explorer/

# 创建文件存储目录
mkdir -p ../../application/backend/files/uploadfiles
mkdir -p ../../application/backend/files/downloadfiles
mkdir -p ../../application/backend/files/images

# 启动区块链浏览器
docker compose -f explorer/docker-compose.yaml up -d
