# registry-consul(kitex注册中心consul的实现)

## docker部署

```shell
docker run -it -d -p 8500:8500 -e CONSUL_BIND_INTERFACE='eth0' \
-v /home/consul/data:/consul/data \
-v /home/consul/config:/consul/config \
 --name=consul1 consul agent -server -bootstrap -ui -client='0.0.0.0'
```

若需要开启ACL(访问控制)

```shell
# 1. 根据示例中docker中挂载的目录
cd /home/consul/data

# 2. 在data文件夹下创建文件 acl.hcl
# 3. 写入如下内容
acl = {
  enabled = true
  default_policy = "deny"
  enable_token_persistence = true 
}
```
UI面板中创建Policies
```shell

# Rules:
service_prefix "" {
   policy = "write"
}

node_prefix "" {
  policy = "read"
}
```
