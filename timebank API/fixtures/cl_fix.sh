#!/bin/bash

# 清除链码容器
echo "清理环境"

mkdir -p channel-artifacts
mkdir -p crypto-config
rm -rf channel-artifacts/*
rm -rf crypto-config/*
sudo docker rm -f $(sudo docker ps -aq)
sudo docker network prune
sudo docker volume prune
echo "清理完毕"
