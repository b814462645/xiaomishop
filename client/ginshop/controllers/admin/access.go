package admin

import (
	"context"
	"ginshop/models"
	"net/http"
	"strings"

	pbAccess "ginshop/proto/rbacAccess"

	"github.com/gin-gonic/gin"
)

// AccessController 结构体定义了访问控制相关的控制器
type AccessController struct {
	BaseController
}

// Index 方法处理访问控制列表页面
func (con AccessController) Index(c *gin.Context) {
	// 创建RBAC访问服务客户端
	rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	// 获取所有访问控制记录
	res, _ := rbacAccessClient.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})

	// 渲染访问控制列表页面
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": res.AccessList,
	})
}

// Add 方法处理添加访问控制页面
func (con AccessController) Add(c *gin.Context) {
	// 获取顶级模块
	rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	res, _ := rbacAccessClient.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})
	// 渲染添加页面
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": res.AccessList,
	})
}

// DoAdd 方法处理添加访问控制的表单提交
func (con AccessController) DoAdd(c *gin.Context) {
	// 获取表单数据
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err1 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")

	// 检查参数是否有效
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	// 创建RBAC访问服务客户端并添加新的访问控制记录
	rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	res, _ := rbacAccessClient.AccessAdd(context.Background(), &pbAccess.AccessAddRequest{
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
	})

	// 处理添加结果
	if !res.Success {
		con.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/access")
}

// Edit 方法处理编辑访问控制页面
func (con AccessController) Edit(c *gin.Context) {
	// 获取要修改的数据ID
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/access")
	}

	// 创建RBAC访问服务客户端
	rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	// 获取指定ID的访问控制记录
	accessRes, _ := rbacAccessClient.AccessGet(context.Background(), &pbAccess.AccessGetRequest{
		Id: int64(id),
	})
	// 获取所有访问控制记录（用于显示顶级模块）
	accessListRes, _ := rbacAccessClient.AccessGet(context.Background(), &pbAccess.AccessGetRequest{})

	// 渲染编辑页面
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     accessRes.AccessList[0],
		"accessList": accessListRes.AccessList,
	})
}

// DoEdit 方法处理编辑访问控制的表单提交
func (con AccessController) DoEdit(c *gin.Context) {
	// 获取表单数据
	id, err1 := models.Int(c.PostForm("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err2 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err3 := models.Int(c.PostForm("module_id"))
	sort, err4 := models.Int(c.PostForm("sort"))
	status, err5 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")

	// 检查参数是否有效
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		con.Error(c, "传入参数错误", "/admin/access")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+models.String(id))
		return
	}

	// 创建RBAC访问服务客户端并更新访问控制记录
	rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
	res, _ := rbacAccessClient.AccessEdit(context.Background(), &pbAccess.AccessEditRequest{
		Id:          int64(id),
		ModuleName:  moduleName,
		Type:        int64(accessType),
		ActionName:  actionName,
		Url:         url,
		ModuleId:    int64(moduleId),
		Sort:        int64(sort),
		Description: description,
		Status:      int64(status),
	})

	// 处理编辑结果
	if !res.Success {
		con.Error(c, "修改数据", "/admin/access/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/access/edit?id="+models.String(id))
	}
}

// Delete 方法处理删除访问控制记录
func (con AccessController) Delete(c *gin.Context) {
	// 获取要删除的记录ID
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access")
	} else {
		// 创建RBAC访问服务客户端并删除指定ID的访问控制记录
		rbacAccessClient := pbAccess.NewRbacAccessService("rbac", models.RbacClient)
		res, _ := rbacAccessClient.AccessDelete(context.Background(), &pbAccess.AccessDeleteRequest{
			Id: int64(id),
		})

		// 处理删除结果
		if res.Success {
			con.Success(c, res.Message, "/admin/access")
		} else {
			con.Error(c, res.Message, "/admin/access")
		}
	}
}
