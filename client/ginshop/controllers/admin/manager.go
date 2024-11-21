package admin

import (
	"context"
	"fmt"
	"ginshop/models"
	pbManager "ginshop/proto/rbacManager"
	pbRole "ginshop/proto/rbacRole"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	minUsernameLength = 2
	minPasswordLength = 6
	maxMobileLength   = 11
	defaultStatus     = 1
)

type ManagerController struct {
	BaseController
}

// Index 显示管理员列表
func (con ManagerController) Index(c *gin.Context) {
	managers, err := con.getManagers()
	if err != nil {
		con.Error(c, "获取管理员列表失败", "/admin")
		return
	}

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managers,
	})
}

// Add 显示添加管理员页面
func (con ManagerController) Add(c *gin.Context) {
	roleList, err := con.getRoleList()
	if err != nil {
		con.Error(c, "获取角色列表失败", "/admin/manager")
		return
	}

	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

// DoAdd 处理添加管理员的请求
func (con ManagerController) DoAdd(c *gin.Context) {
	manager, err := con.getManagerData(c)
	if err != nil {
		con.Error(c, err.Error(), "/admin/manager/add")
		return
	}

	if err := con.validateManagerData(manager.Username, manager.Password, manager.Mobile); err != nil {
		con.Error(c, err.Error(), "/admin/manager/add")
		return
	}

	if exists, err := con.checkManagerExists(manager.Username); err != nil {
		con.Error(c, "检查管理员是否存在时出错", "/admin/manager/add")
		return
	} else if exists {
		con.Error(c, "此管理员已存在", "/admin/manager/add")
		return
	}

	if err := con.addManager(manager); err != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}

	con.Success(c, "增加管理员成功", "/admin/manager")
}

// Edit 显示编辑管理员页面
func (con ManagerController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}

	manager, err := con.getManager(id)
	if err != nil {
		con.Error(c, "获取管理员信息失败", "/admin/manager")
		return
	}

	roleList, err := con.getRoleList()
	if err != nil {
		con.Error(c, "获取角色列表失败", "/admin/manager")
		return
	}

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

// DoEdit 处理编辑管理员的请求
func (con ManagerController) DoEdit(c *gin.Context) {
	manager, err := con.getEditManagerData(c)
	if err != nil {
		con.Error(c, err.Error(), "/admin/manager")
		return
	}

	if err := con.validateEditManagerData(manager.Password, manager.Mobile); err != nil {
		con.Error(c, err.Error(), fmt.Sprintf("/admin/manager/edit?id=%d", manager.Id))
		return
	}

	if err := con.editManager(manager); err != nil {
		con.Error(c, "修改数据失败", fmt.Sprintf("/admin/manager/edit?id=%d", manager.Id))
		return
	}

	con.Success(c, "修改数据成功", "/admin/manager")
}

// Delete 处理删除管理员的请求
func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager")
		return
	}

	if err := con.deleteManager(id); err != nil {
		con.Error(c, "删除数据失败", "/admin/manager")
		return
	}

	con.Success(c, "删除数据成功", "/admin/manager")
}

// 辅助方法

func (con ManagerController) getManagers() ([]*pbManager.ManagerModel, error) {
	rbacManagerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	res, err := rbacManagerClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{})
	if err != nil {
		return nil, err
	}
	return res.ManagerList, nil
}

func (con ManagerController) getRoleList() ([]*pbRole.RoleModel, error) {
	rbacRoleClient := pbRole.NewRbacRoleService("rbac", models.RbacClient)
	res, err := rbacRoleClient.RoleGet(context.Background(), &pbRole.RoleGetRequest{})
	if err != nil {
		return nil, err
	}
	return res.RoleList, nil
}

func (con ManagerController) getManagerData(c *gin.Context) (*pbManager.ManagerAddRequest, error) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		return nil, fmt.Errorf("角色ID格式错误")
	}
	return &pbManager.ManagerAddRequest{
		Username: strings.TrimSpace(c.PostForm("username")),
		Password: strings.TrimSpace(c.PostForm("password")),
		Email:    strings.TrimSpace(c.PostForm("email")),
		Mobile:   strings.TrimSpace(c.PostForm("mobile")),
		Status:   defaultStatus,
		RoleId:   int64(roleId),
		AddTime:  int64(models.GetUnix()),
	}, nil
}

func (con ManagerController) validateManagerData(username, password, mobile string) error {
	if len(username) < minUsernameLength || len(password) < minPasswordLength {
		return fmt.Errorf("用户名或者密码的长度不合法")
	}
	if len(mobile) > maxMobileLength {
		return fmt.Errorf("手机号码长度不合法")
	}
	return nil
}

func (con ManagerController) checkManagerExists(username string) (bool, error) {
	rbacManagerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	resGet, err := rbacManagerClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Username: username,
	})
	if err != nil {
		return false, err
	}
	return len(resGet.ManagerList) > 0, nil
}

func (con ManagerController) addManager(manager *pbManager.ManagerAddRequest) error {
	rbacManagerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	manager.Password = models.Md5(manager.Password)
	resAdd, err := rbacManagerClient.ManagerAdd(context.Background(), manager)
	if err != nil {
		return err
	}
	if !resAdd.Success {
		return fmt.Errorf("增加管理员失败: %s", resAdd.Message)
	}
	return nil
}

func (con ManagerController) getManager(id int) (*pbManager.ManagerModel, error) {
	rbacManagerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	managerRes, err := rbacManagerClient.ManagerGet(context.Background(), &pbManager.ManagerGetRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}
	if len(managerRes.ManagerList) == 0 {
		return nil, fmt.Errorf("管理员不存在")
	}
	return managerRes.ManagerList[0], nil
}

func (con ManagerController) getEditManagerData(c *gin.Context) (*pbManager.ManagerEditRequest, error) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		return nil, fmt.Errorf("ID格式错误")
	}
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		return nil, fmt.Errorf("角色ID格式错误")
	}
	return &pbManager.ManagerEditRequest{
		Id:       int64(id),
		Username: strings.TrimSpace(c.PostForm("username")),
		Password: strings.TrimSpace(c.PostForm("password")),
		Email:    strings.TrimSpace(c.PostForm("email")),
		Mobile:   strings.TrimSpace(c.PostForm("mobile")),
		RoleId:   int64(roleId),
	}, nil
}

func (con ManagerController) validateEditManagerData(password, mobile string) error {
	if password != "" && len(password) < minPasswordLength {
		return fmt.Errorf("密码的长度不合法 密码长度不能小于6位")
	}
	if len(mobile) > maxMobileLength {
		return fmt.Errorf("手机号码长度不合法")
	}
	return nil
}

func (con ManagerController) editManager(manager *pbManager.ManagerEditRequest) error {
	if manager.Password != "" {
		manager.Password = models.Md5(manager.Password)
	}
	managerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	res, err := managerClient.ManagerEdit(context.Background(), manager)
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("修改数据失败: %s", res.Message)
	}
	return nil
}

func (con ManagerController) deleteManager(id int) error {
	managerClient := pbManager.NewRbacManagerService("rbac", models.RbacClient)
	res, err := managerClient.ManagerDelete(context.Background(), &pbManager.ManagerDeleteRequest{
		Id: int64(id),
	})
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("删除数据失败: %s", res.Message)
	}
	return nil
}
