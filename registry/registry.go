package registry

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
	"github.com/towelong/registry-consul/consul"
	"net"
	"strconv"
)

var _ registry.Registry = (*consulRegistry)(nil)

type consulRegistry struct {
	cli *api.Client
}

func NewConsulRegistry(address string) (registry.Registry, error) {
	cli, err := consul.NewConsulClient(address)
	if err != nil {
		return nil, err
	}
	c := &consulRegistry{
		cli: cli,
	}
	return c, nil
}

func NewConsulRegisterWithConfig(config *api.Config) (*consulRegistry, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &consulRegistry{cli: client}, nil
}

func (c *consulRegistry) Register(info *registry.Info) error {
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return fmt.Errorf("parse registry info addr error: %w", err)
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("parse registry info port error: %w", err)
	}
	tags := make([]string, 0)
	for k, v := range info.Tags {
		tags = append(tags, fmt.Sprintf("%s=%s", k, v))
	}
	asr := &api.AgentServiceRegistration{
		ID:      info.ServiceName,
		Name:    info.ServiceName,
		Tags:    tags,
		Address: host,
		Port:    p,
		Weights: &api.AgentWeights{
			Passing: info.Weight,
		},
	}
	return c.cli.Agent().ServiceRegister(asr)
}

func (c *consulRegistry) Deregister(info *registry.Info) error {
	return c.cli.Agent().ServiceDeregister(info.ServiceName)
}
