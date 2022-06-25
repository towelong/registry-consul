package resolver

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hashicorp/consul/api"
	"github.com/towelong/registry-consul/consul"
	"strings"
)

var _ discovery.Resolver = (*consulResolver)(nil)

type consulResolver struct {
	cli *api.Client
}

func (c *consulResolver) Target(_ context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

func (c *consulResolver) Resolve(_ context.Context, desc string) (discovery.Result, error) {
	services, _, err := c.cli.Catalog().Service(desc, "", nil)
	if err != nil {
		return discovery.Result{}, fmt.Errorf("get services error: %w", err)
	}
	instances := make([]discovery.Instance, 0, len(services))
	for _, srv := range services {
		tagMap := make(map[string]string)
		for _, tag := range srv.ServiceTags {
			t := strings.Split(tag, "=")
			tagMap[t[0]] = t[1]
		}
		instances = append(instances, discovery.NewInstance(
			"tcp",
			fmt.Sprintf("%s:%d", srv.ServiceAddress, srv.ServicePort),
			srv.ServiceWeights.Passing,
			tagMap,
		))
	}
	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}, nil
}

func (c *consulResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (c *consulResolver) Name() string {
	return "consul"
}

func NewConsulResolver(address string) (discovery.Resolver, error) {
	cli, err := consul.NewConsulClient(address)
	if err != nil {
		return nil, err
	}
	return &consulResolver{cli: cli}, nil
}

func NewConsulResolverConfig(config *api.Config) (discovery.Resolver, error) {
	cli, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &consulResolver{cli: cli}, nil
}
