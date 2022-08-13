package resolver

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/stretchr/testify/require"
	registry2 "github.com/towelong/registry-consul/registry"
	"testing"
)

const (
	serviceName  = "API"
	registryAddr = "127.0.0.1:8500"
	addr         = "127.0.0.1:8001"
)

func initRegistry() {
	rg, err := registry2.NewConsulRegistry(registryAddr)
	if err != nil {
		panic(err)
	}
	info := registry.Info{
		ServiceName: serviceName,
		Addr:        utils.NewNetAddr("tcp", addr),
		Weight:      99,
		Tags: map[string]string{
			"version": "1.0.0",
		},
	}
	err = rg.Register(&info)
	if err != nil {
		panic(err)
	}
}

func TestNewConsulResolver(t *testing.T) {
	rs, err := NewConsulResolver(registryAddr)
	require.Nil(t, err)

	initRegistry()
	desc := rs.Target(context.TODO(), rpcinfo.NewEndpointInfo(serviceName, "", nil, nil))
	require.Equal(t, serviceName, desc)

	res, err := rs.Resolve(context.TODO(), desc)
	require.Nil(t, err)
	fmt.Printf("%v", res.Instances[0].Address())
}
