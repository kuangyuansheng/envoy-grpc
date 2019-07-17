# envoyproxy grpc测试实例
envoy微服务 服务发现,服务治理demo

### 安装
```
#首先要安装golang, 然后设置GOPATH
git clone github.com/kuangyuansheng/envoy-grpc $GOPATH/src/envoy-grpc

# grpc健康检查
go get -u google.golang.org/grpc/health
```

### 期望
rpcCli访问19000端口,  envoy会自动路由到18000,17000,  当关掉其中一个服务, 通过健康检查, envoy仍能路由到健康的节点

### 运行
```
# 运行envoy
./envoy -c envoy.yaml

# 开启微服务
make
./rpcSrv -p 18000
./rpcSrv -p 17000
# 服务开启一两分钟后会持续输出心跳
# 2019/07/11 10:35:28 service started
# 2019/07/11 10:35:28 checking............Watch
# 2019/07/11 10:35:29 checking............Watch
# 2019/07/11 10:35:30 checking............Watch
# 2019/07/11 10:35:31 checking............Watch
# 2019/07/11 10:35:32 checking............Watch
# 2019/07/11 10:35:33 checking............Watch

# 运行客户端
./rpcCli -a 127.0.0.1:19000
```




