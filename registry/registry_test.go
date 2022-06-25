package registry

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	addr         = "127.0.0.1:8001"
	serviceName  = "API"
	registryAddr = "127.0.0.1:8500"
)

func initRegistry() registry.Registry {
	rg, err := NewConsulRegistry(registryAddr)
	if err != nil {
		panic(err)
	}
	return rg
}

func TestNewConsulRegistry(t *testing.T) {
	_, err := NewConsulRegistry(registryAddr)
	require.Nil(t, err)
}

func Test_consulRegistry_Register(t *testing.T) {
	rg := initRegistry()
	info := registry.Info{
		ServiceName: serviceName,
		Addr:        utils.NewNetAddr("tcp", addr),
		Weight:      99,
		Tags: map[string]string{
			"version": "1.0.0",
		},
	}
	err := rg.Register(&info)
	require.Nil(t, err)
}

func Test_consulRegistry_Deregister(t *testing.T) {
	rg := initRegistry()
	err := rg.Deregister(&registry.Info{
		ServiceName: serviceName,
	})
	require.Nil(t, err)
}
