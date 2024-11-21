package models

import (
	"github.com/go-micro/plugins/v4/registry/kubernetes"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var RbacClient client.Client

func init() {
	k8sReg := kubernetes.NewRegistry(
		registry.Addrs("124.220.36.152:6443"),
	)

	srv := micro.NewService(
		micro.Name("rbac.client"),
		micro.Registry(k8sReg),
	)

	srv.Init()
	RbacClient = srv.Client()
}
