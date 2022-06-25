package consul

import "github.com/hashicorp/consul/api"

func NewConsulClient(address string) (*api.Client, error) {
	client, err := api.NewClient(&api.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
