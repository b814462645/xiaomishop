package main

import (
	"captcha/handler"
	pb "captcha/proto/captcha"

	"github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	//"go-micro.dev/v4/cmd"
)

var (
	service = "captcha"
	version = "latest"
)

func main() {
	// consulReg := consul.NewRegistry(
	// 	registry.Addrs("127.0.0.1:8500"),
	// )
	// // Create service
	// srv := micro.NewService(
	// 	micro.Registry(consulReg),
	// )
	// srv.Init(
	// 	micro.Name(service),
	// 	micro.Version(version),
	// )

	// 创建一个微服务，使用Kubernetes作为服务注册中心
	srv := micro.NewService(
		micro.Name("captcha"),
		micro.Registry(kubernetes.NewRegistry()), // 使用 Kubernetes 注册表
	)
	srv.Init()

	// Register handler
	if err := pb.RegisterCaptchaHandler(srv.Server(), new(handler.Captcha)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
