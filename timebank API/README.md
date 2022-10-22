````
将`GOPATH`设置为`/root/go`,将timebank项目放到`$GOPATH/src/`下
````
`sudo vim /etc/profile`
````
配置生效
````
`source /etc/profile`
````
将`fabric-samples`的`bin`目录里的二进制文件全都复制在`$GOPSTH/bin/`下
````
在`/etc/hosts`中添加：
````
`127.0.0.1  orderer.jxnu.com`
`127.0.0.1  peer0.jxnu.qsh.com`
`127.0.0.1  peer0.jxnu.yh.com`
`127.0.0.1  peer0.jxnu.gqc.com`
````
启动配置
````
`cd application && cd service && ./start.sh`s
````
清理配置
````
`cd fixtures&&./cl_fix.sh`
