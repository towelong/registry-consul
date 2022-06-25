package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/client"
	consul "github.com/towelong/registry-consul/resolver"
	"log"
	"time"
)

func main() {
	// 未开启ACL直接填入consul注册中心地址
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	// 开启了ACL访问控制则使用如下代码
	//r, err := consul.NewConsulResolverConfig(&api2.Config{
	//	Address: "127.0.0.1:8500",
	//	Token:   "0bdbbb87-xxxx-ada3-c4a8-a7c0fb095867",
	//})
	if err != nil {
		log.Fatal(err)
	}
	cli := hello.MustNewClient(
		"Hello",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)
	for {
		resp, err := cli.Echo(context.Background(), &api.Request{Message: "Hello"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
