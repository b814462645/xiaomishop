package handler

import (
	"context"

	"rbac/models"
	pb "rbac/proto/rbacManager"
)

type RbacManager struct{}

func (e *RbacManager) ManagerGet(ctx context.Context, req *pb.ManagerGetRequest, rsp *pb.ManagerGetResponse) error {
	managerList := []models.Manager{}
	query := models.DB.Where("1=1")
	if req.Id > 0 {
		query = query.Where("id = ?", req.Id)
	}
	if len(req.Username) > 0 {
		query = query.Where("username = ?", req.Username)
	}
	query.Preload("Role").Find(&managerList)

	var tempList []*pb.ManagerModel
	for _, v := range managerList {
		tempList = append(tempList, &pb.ManagerModel{
			Id:       int64(v.Id),
			Username: v.Username,
			Mobile:   v.Mobile,
			Email:    v.Email,
			Status:   int64(v.Status),
			RoleId:   int64(v.RoleId),
			AddTime:  int64(v.AddTime),
			IsSuper:  int64(v.IsSuper),
			Role: &pb.RoleModel{
				Title:       v.Role.Title,
				Description: v.Role.Description,
			},
		})
	}

	rsp.ManagerList = tempList
	return nil
}

func (e *RbacManager) ManagerAdd(ctx context.Context, req *pb.ManagerAddRequest, rsp *pb.ManagerAddResponse) error {
	manager := models.Manager{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		RoleId:   int(req.RoleId),
		Status:   int(req.Status),
		AddTime:  int(req.AddTime),
	}
	err := models.DB.Create(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = "增加数据失败"
	} else {
		rsp.Success = true
		rsp.Message = "增加数据成功"
	}
	return err
}

func (e *RbacManager) ManagerEdit(ctx context.Context, req *pb.ManagerEditRequest, rsp *pb.ManagerEditResponse) error {
	manager := models.Manager{Id: int(req.Id)}
	models.DB.Find(&manager)
	manager.Username = req.Username
	manager.Email = req.Email
	manager.Mobile = req.Mobile
	manager.RoleId = int(req.RoleId)

	if req.Password != "" {
		manager.Password = req.Password
	}
	err := models.DB.Save(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = "修改数据失败"
	} else {
		rsp.Success = true
		rsp.Message = "修改数据成功"
	}
	return err
}

func (e *RbacManager) ManagerDelete(ctx context.Context, req *pb.ManagerDeleteRequest, rsp *pb.ManagerDeleteResponse) error {
	manager := models.Manager{Id: int(req.Id)}
	err := models.DB.Delete(&manager).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = "删除数据失败"
	} else {
		rsp.Success = true
		rsp.Message = "删除数据成功"
	}

	return err
}
