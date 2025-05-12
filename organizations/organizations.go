package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/tempest-org/tempest/organizations/internal/config"
	"github.com/tempest-org/tempest/organizations/internal/server"
	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"
	"github.com/tempest-org/tempest/pkg/datasource"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/organizations.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)


	res, err := datasource.Migrate(context.Background(), c.Database.URI, "./migrations")
	if err != nil {
		panic(err)
	}
	logx.Info(res)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		organizations.RegisterOrganizationsServer(grpcServer, server.NewOrganizationsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
