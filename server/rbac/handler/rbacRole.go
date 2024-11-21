package handler

import (
	"context"
	"strconv"

	"rbac/models"
	pb "rbac/proto/rbacRole"

	"gorm.io/gorm"
)

type RbacRole struct{}

func (e *RbacRole) RoleGet(ctx context.Context, req *pb.RoleGetRequest, rsp *pb.RoleGetResponse) error {
	roleList := []models.Role{}
	where := "1=1"
	if req.Id > 0 {
		where += " AND id = " + strconv.Itoa(int(req.Id))
	}
	err := models.DB.Where(where).Find(&roleList).Error
	if err != nil {
		return err
	}
	var tempList []*pb.RoleModel
	for _, v := range roleList {
		tempList = append(tempList, &pb.RoleModel{
			Id:          int64(v.Id),
			Title:       v.Title,
			Description: v.Description,
			Status:      int64(v.Status),
			AddTime:     int64(v.AddTime),
		})
	}
	rsp.RoleList = tempList
	return nil
}

func (e *RbacRole) RoleAdd(ctx context.Context, req *pb.RoleAddRequest, rsp *pb.RoleAddResponse) error {
	role := models.Role{
		Title:       req.Title,
		Description: req.Description,
		AddTime:     int(req.AddTime),
		Status:      int(req.Status),
	}
	err := models.DB.Create(&role).Error
	if err != nil {
		rsp.Message = "增加角色失败 请重试"
		rsp.Success = false
	} else {
		rsp.Message = "增加成功"
		rsp.Success = true
	}
	return err
}

func (e *RbacRole) RoleEdit(ctx context.Context, req *pb.RoleEditRequest, rsp *pb.RoleEditResponse) error {
	role := models.Role{Id: int(req.Id)}
	models.DB.Find(&role)
	role.Title = req.Title
	role.Description = req.Description

	err := models.DB.Save(&role).Error
	if err != nil {
		rsp.Message = "修改数据失败"
		rsp.Success = false
	} else {
		rsp.Message = "修改数据成功"
		rsp.Success = true
	}
	return err
}

func (e *RbacRole) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest, rsp *pb.RoleDeleteResponse) error {
	role := models.Role{Id: int(req.Id)}
	err := models.DB.Delete(&role).Error

	if err != nil {
		rsp.Message = "删除数据失败"
		rsp.Success = false
	} else {
		rsp.Message = "删除数据成功"
		rsp.Success = true
	}

	return err
}

func (e *RbacRole) RoleAuth(ctx context.Context, req *pb.RoleAuthRequest, rsp *pb.RoleAuthResponse) error {
	//2、获取所有的权限
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Order("sort DESC").Find(&accessList)

	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", req.RoleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[int(accessList[i].Id)]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[int(accessList[i].AccessItem[j].Id)]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

	var tempList []*pb.AccessModel
	for _, v := range accessList {
		var tempItemList []*pb.AccessModel
		for _, k := range v.AccessItem {
			tempItemList = append(tempItemList, &pb.AccessModel{
				Id:          int64(k.Id),
				ModuleName:  k.ModuleName,
				ActionName:  k.ActionName,
				Type:        int64(k.Type),
				Url:         k.Url,
				ModuleId:    int64(k.ModuleId),
				Sort:        int64(k.Sort),
				Description: k.Description,
				Status:      int64(k.Status),
				Checked:     k.Checked,
				AddTime:     int64(k.AddTime),
			})
		}
		tempList = append(tempList, &pb.AccessModel{
			Id:          int64(v.Id),
			ModuleName:  v.ModuleName,
			ActionName:  v.ActionName,
			Type:        int64(v.Type),
			Url:         v.Url,
			ModuleId:    int64(v.ModuleId),
			Sort:        int64(v.Sort),
			Description: v.Description,
			Status:      int64(v.Status),
			AddTime:     int64(v.AddTime),
			Checked:     v.Checked,
			AccessItem:  tempItemList,
		})
	}

	rsp.AccessList = tempList
	return nil

}

func (e *RbacRole) RoleDoAuth(ctx context.Context, req *pb.RoleDoAuthRequest, rsp *pb.RoleDoAuthResponse) error {
	//删除当前角色对应的权限
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", req.RoleId).Delete(&roleAccess)

	//增加当前角色对应的权限
	for _, v := range req.AccessIds {
		roleAccess.RoleId = int(req.RoleId)
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}

	rsp.Success = true
	rsp.Message = "授权成功"
	return nil
}

func (e *RbacRole) RoleMiddlewaresAuth(ctx context.Context, req *pb.RoleMiddlewaresAuthRequest, rsp *pb.RoleMiddlewaresAuthResponse) error {
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", req.RoleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}
	// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
	// pathname      /admin/manager
	access := models.Access{}
	models.DB.Where("url = ?", req.UrlPath).Find(&access)
	//3、判断当前访问的url对应的权限id 是否在权限列表的id中
	if _, ok := roleAccessMap[access.Id]; !ok {
		rsp.HasPermission = false
	} else {
		rsp.HasPermission = true
	}
	return nil
}
