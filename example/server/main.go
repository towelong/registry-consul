package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/towelong/registry-consul/registry"
	"log"
	"net"
)

type HelloImpl struct{}

func (h *HelloImpl) Echo(_ context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{
		Message: req.Message,
	}
	return
}

func main() {
	// 未开启ACL直接填入consul注册中心地址
	r, err := registry.NewConsulRegistry("127.0.0.1:8500")
	// 开启了ACL访问控制则使用如下代码
	//r, err := consul.NewConsulRegistryConfig(&api2.Config{
	//	Address: "127.0.0.1:8500",
	//	Token:   "0bdbbb87-xxxx-ada3-c4a8-a7c0fb095867",
	//})
	if err != nil {
		panic(err)
	}
	svr := hello.NewServer(
		new(HelloImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Hello"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
