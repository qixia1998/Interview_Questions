package main

import (
	"flag"
	"fmt"

	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/courseware"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/internal/config"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/internal/server"
	"Interview_Questions/Go-Zero/go-zero-courseware/courseware/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/courseware.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		courseware.RegisterCoursewareServer(grpcServer, server.NewCoursewareServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
