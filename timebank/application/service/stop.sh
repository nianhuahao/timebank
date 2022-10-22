echo "开始清理"
cd ..
cd ..
cd fixtures
./cl_fix.sh
echo "清理完成"
docker ps
echo "查看端口ip"
lsof -i :8410
