package main

import (
	"rbac/handler"
	"rbac/models"
	pbAccess "rbac/proto/rbacAccess"
	pbLogin "rbac/proto/rbacLogin"
	pbManager "rbac/proto/rbacManager"
	pbRole "rbac/proto/rbacRole"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "rbac"
	version = "latest"
)

func main() {

	addr := models.Config.Section("consul").Key("addr").String()
	consulReg := consul.NewRegistry(
		registry.Addrs(addr),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pbLogin.RegisterRbacLoginHandler(srv.Server(), new(handler.RbacLogin)); err != nil {
		logger.Fatal(err)
	}
	if err := pbRole.RegisterRbacRoleHandler(srv.Server(), new(handler.RbacRole)); err != nil {
		logger.Fatal(err)
	}
	if err := pbManager.RegisterRbacManagerHandler(srv.Server(), new(handler.RbacManager)); err != nil {
		logger.Fatal(err)
	}
	if err := pbAccess.RegisterRbacAccessHandler(srv.Server(), new(handler.RbacAccess)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
