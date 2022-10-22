#!/bin/bash
echo "一、清理环境"
./cl_fix.sh

echo "二、生成证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
cryptogen generate --config=./crypto-config.yaml

echo "三、创建排序通道创世区块"
configtxgen -profile FourOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID firstchannel

echo "四、生成通道配置事务'mychannel.tx'"
configtxgen -profile FourOrgsChannel -outputCreateChannelTx ./channel-artifacts/mychannel.tx -channelID mychannel

echo "五、为 Qsh 定义锚节点"
configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/QshAnchor.tx -channelID mychannel -asOrg Qsh

echo "六、为 Yh 定义锚节点"
configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/YhAnchor.tx -channelID mychannel -asOrg Yh

echo "六、为 Yh 定义锚节点"
configtxgen -profile FourOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/GqcAnchor.tx -channelID mychannel -asOrg Gqc

 echo "区块链 : 启动"
 docker-compose up -d
 echo "正在等待节点的启动完成,等待3秒"
 sleep 3
 echo "检查容器"
 docker ps