package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"tweets/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer,NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}


//protoc newsfeed.proto  --proto_path=./third_party --proto_path=./api/newsfeeds/v1/  --proto_path=./api/tweet/v1/ --go_out=paths=source_relative:.