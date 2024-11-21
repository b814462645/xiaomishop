package admin

import (
	"context"
	"fmt"
	"ginshop/models"
	pbRole "ginshop/proto/rbacRole"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	roleTimeoutDuration = 5 * time.Second
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacClient.RoleGet(ctx, &pbRole.RoleGetRequest{})
	if err != nil {
		con.Error(c, fmt.Sprintf("获取角色列表失败: %v", err), "/admin")
		return
	}

	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": res.RoleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.TrimSpace(c.PostForm("title"))
	description := strings.TrimSpace(c.PostForm("description"))

	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacClient.RoleAdd(ctx, &pbRole.RoleAddRequest{
		Title:       title,
		Description: description,
		AddTime:     models.GetUnix(),
		Status:      1,
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("增加角色失败: %v", err), "/admin/role/add")
		return
	}

	if !res.Success {
		con.Error(c, fmt.Sprintf("增加角色失败: %s", res.Message), "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacClient.RoleGet(ctx, &pbRole.RoleGetRequest{
		Id: int64(id),
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("获取角色信息失败: %v", err), "/admin/role")
		return
	}

	if len(res.RoleList) == 0 {
		con.Error(c, "角色不存在", "/admin/role")
		return
	}

	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": res.RoleList[0],
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	title := strings.TrimSpace(c.PostForm("title"))
	description := strings.TrimSpace(c.PostForm("description"))

	if title == "" {
		con.Error(c, "角色的标题不能为空", fmt.Sprintf("/admin/role/edit?id=%d", id))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacClient.RoleEdit(ctx, &pbRole.RoleEditRequest{
		Id:          int64(id),
		Title:       title,
		Description: description,
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("修改数据失败: %v", err), fmt.Sprintf("/admin/role/edit?id=%d", id))
		return
	}

	if !res.Success {
		con.Error(c, fmt.Sprintf("修改数据失败: %s", res.Message), fmt.Sprintf("/admin/role/edit?id=%d", id))
	} else {
		con.Success(c, "修改数据成功", fmt.Sprintf("/admin/role/edit?id=%d", id))
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacClient.RoleDelete(ctx, &pbRole.RoleDeleteRequest{
		Id: int64(id),
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("删除数据失败: %v", err), "/admin/role")
		return
	}

	if res.Success {
		con.Success(c, "删除数据成功", "/admin/role")
	} else {
		con.Error(c, fmt.Sprintf("删除数据失败: %s", res.Message), "/admin/role")
	}
}

func (con RoleController) Auth(c *gin.Context) {
	roleId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacRoleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacRoleClient.RoleAuth(ctx, &pbRole.RoleAuthRequest{
		RoleId: int64(roleId),
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("获取权限列表失败: %v", err), "/admin/role")
		return
	}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": res.AccessList,
	})
}

func (con RoleController) DoAuth(c *gin.Context) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}

	accessIds := c.PostFormArray("access_node[]")

	ctx, cancel := context.WithTimeout(c.Request.Context(), roleTimeoutDuration)
	defer cancel()

	rbacRoleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacRoleClient.RoleDoAuth(ctx, &pbRole.RoleDoAuthRequest{
		RoleId:    int64(roleId),
		AccessIds: accessIds,
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("授权失败: %v", err), fmt.Sprintf("/admin/role/auth?id=%d", roleId))
		return
	}

	if res.Success {
		con.Success(c, "授权成功", fmt.Sprintf("/admin/role/auth?id=%d", roleId))
	} else {
		con.Error(c, fmt.Sprintf("授权失败: %s", res.Message), fmt.Sprintf("/admin/role/auth?id=%d", roleId))
	}
}
