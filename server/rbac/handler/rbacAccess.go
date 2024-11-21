package handler

import (
	"context"
	"strconv"

	"rbac/models"
	pb "rbac/proto/rbacAccess"
)

type RbacAccess struct{}

func (e *RbacAccess) AccessGet(ctx context.Context, req *pb.AccessGetRequest, rsp *pb.AccessGetResponse) error {
	accessList := []models.Access{}
	where := "1=1"
	if req.Id > 0 {
		where += " AND id=" + strconv.Itoa(int(req.Id))
	} else {
		where += " AND module_id = 0"
	}
	models.DB.Where(where).Preload("AccessItem").Find(&accessList)

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
			AccessItem:  tempItemList,
		})
	}
	rsp.AccessList = tempList
	return nil
}

func (e *RbacAccess) AccessAdd(ctx context.Context, req *pb.AccessAddRequest, rsp *pb.AccessAddResponse) error {
	access := models.Access{
		ModuleName:  req.ModuleName,
		Type:        int(req.Type),
		ActionName:  req.ActionName,
		Url:         req.Url,
		ModuleId:    int(req.ModuleId),
		Sort:        int(req.Sort),
		Description: req.Description,
		Status:      int(req.Status),
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		rsp.Success = false
		rsp.Message = "添加数据失败"
	} else {
		rsp.Success = true
		rsp.Message = "添加数据成功"
	}

	return err
}

func (e *RbacAccess) AccessEdit(ctx context.Context, req *pb.AccessEditRequest, rsp *pb.AccessEditResponse) error {
	access := models.Access{Id: int(req.Id)}
	models.DB.Find(&access)
	access.ModuleName = req.ModuleName
	access.Type = int(req.Type)
	access.ActionName = req.ActionName
	access.Url = req.Url
	access.ModuleId = int(req.ModuleId)
	access.Sort = int(req.Sort)
	access.Description = req.Description
	access.Status = int(req.Status)

	err := models.DB.Save(&access).Error

	if err != nil {
		rsp.Success = false
		rsp.Message = "编辑数据失败"
	} else {
		rsp.Success = true
		rsp.Message = "编辑数据成功"
	}
	return err
}

func (e *RbacAccess) AccessDelete(ctx context.Context, req *pb.AccessDeleteRequest, rsp *pb.AccessDeleteResponse) error {
	//获取我们要删除的数据
	access := models.Access{Id: int(req.Id)}
	models.DB.Find(&access)
	if access.ModuleId == 0 { //顶级模块
		accessList := []models.Access{}
		models.DB.Where("module_id = ?", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			rsp.Success = false
			rsp.Message = "当前模块下面有菜单或者操作，请删除菜单或者操作以后再来删除这个数据"
		} else {
			models.DB.Delete(&access)
			rsp.Success = true
			rsp.Message = "删除数据成功"
		}
	} else { //操作 或者菜单
		models.DB.Delete(&access)
		rsp.Success = true
		rsp.Message = "删除数据成功"
	}
	return nil
}
