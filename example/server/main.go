package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/towelong/registry-consul/registry"
	"net"
	"os"
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
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	svr := hello.NewServer(
		new(HelloImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Hello"}),
		server.WithServiceAddr(addr),
	)
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//klog.SetOutput(f)
	klog.SetLevel(klog.LevelDebug)
	klog.Debug("Debug中 ----")
	klog.CtxInfof(context.Background(), "%s", "ctx 666")
	klog.Error("服务器开小差了~")
	if err := svr.Run(); err != nil {
		klog.Error("server stopped with error:", err)
	} else {
		klog.Info("server stopped")
	}
}
