package handler

import (
	"context"
	"fmt"

	"rbac/models"
	pb "rbac/proto/rbacLogin"
)

type RbacLogin struct{}

func (e *RbacLogin) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	fmt.Println(req)
	managerList := []models.Manager{}
	err := models.DB.Where("username=? AND password=?", req.Username, req.Password).Find(&managerList).Error

	var tempList []*pb.ManagerModel
	for i := 0; i < len(managerList); i++ {
		tempList = append(tempList, &pb.ManagerModel{
			Id:       int64(managerList[i].Id),
			Username: managerList[i].Username,
			Password: managerList[i].Password,
			Mobile:   managerList[i].Mobile,
			Email:    managerList[i].Email,
			Status:   int64(managerList[i].Status),
			RoleId:   int64(managerList[i].RoleId),
			AddTime:  int64(managerList[i].AddTime),
			IsSuper:  int64(managerList[i].IsSuper),
		})
	}

	if len(managerList) > 0 {
		rsp.IsLogin = true
	} else {
		rsp.IsLogin = false
	}
	rsp.Userlist = tempList

	return err
}
