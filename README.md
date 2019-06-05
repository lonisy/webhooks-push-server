# Webhooks 代码 push 服务端程序
50行代码实现 Webhooks 代码 push 服务端程序。



## 构建
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webhook-push-server main.go

```
## 运行方式
```
nohup ./webhook-push-server -log_dir=./日志目录 &

```

## 仓库 Webhooks 调用

```
curl -v http://{youwebhookshost}:4433/?git={yourProjectName}
```

## 服务端配置

脚本参考: project-demo-update.sh

```
sudo touch /home/{yourProjectName}-update.sh 
sudo chmod 777 /home/{yourProjectName}-update.sh
```
