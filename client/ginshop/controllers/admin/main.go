package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"ginshop/models"
	pbRole "ginshop/proto/rbacRole"
	"net/http"

	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// MainController 处理主要的管理功能
type MainController struct {
	BaseController
}

// Index 处理管理员首页
func (mc MainController) Index(c *gin.Context) {
	userinfo, err := mc.getUserInfo(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	// 调用RBAC服务获取用户权限
	rbacRoleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	resAccess, err := rbacRoleClient.RoleAuth(context.Background(), &pbRole.RoleAuthRequest{
		RoleId: int64(userinfo.RoleId),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取权限失败"})
		return
	}

	// 渲染首页
	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
		"username":   userinfo.Username,
		"accessList": resAccess.AccessList,
		"isSuper":    userinfo.IsSuper,
	})
}

// getUserInfo 从session中获取用户信息
func (mc MainController) getUserInfo(c *gin.Context) (*models.Manager, error) {
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	userinfoStr, ok := userinfo.(string)
	if !ok {
		return nil, errors.New("未登录")
	}

	var userinfoStruct []models.Manager
	if err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct); err != nil {
		return nil, err
	}

	if len(userinfoStruct) == 0 {
		return nil, errors.New("用户信息为空")
	}

	return &userinfoStruct[0], nil
}

// Welcome 渲染欢迎页面
func (mc MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// ChangeStatus 通用的状态修改方法
func (mc MainController) ChangeStatus(c *gin.Context) {
	id, table, field, err := mc.getChangeParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	sql := fmt.Sprintf("UPDATE %s SET %s = ABS(%s - 1) WHERE id = ?", table, field, field)
	if err = models.DB.Exec(sql, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "修改失败，请重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "修改成功"})
}

// ChangeNum 通用的数值修改方法
func (mc MainController) ChangeNum(c *gin.Context) {
	id, table, field, err := mc.getChangeParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	num, err := models.Int(c.Query("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "数值参数错误"})
		return
	}

	sql := fmt.Sprintf("UPDATE %s SET %s = %d WHERE id = ?", table, field, num)
	if err = models.DB.Exec(sql, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "修改数据失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "修改成功"})
}

// getChangeParams 获取修改操作的公共参数
func (mc MainController) getChangeParams(c *gin.Context) (int, string, string, error) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		return 0, "", "", errors.New("ID参数错误")
	}

	table := c.Query("table")
	field := c.Query("field")
	if table == "" || field == "" {
		return 0, "", "", errors.New("表名或字段名不能为空")
	}

	return id, table, field, nil
}

// FlushAll 清除Redis缓存
func (mc MainController) FlushAll(c *gin.Context) {
	err := models.CacheDb.FlushAll()
	if err != nil {
		mc.Error(c, "清除Redis缓存失败", "/admin")
		return
	}
	mc.Success(c, "清除Redis缓存数据成功", "/admin")
}
