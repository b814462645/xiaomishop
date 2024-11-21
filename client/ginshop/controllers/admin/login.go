package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"ginshop/models"
	pbRbac "ginshop/proto/rbacLogin"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	timeoutDuration = 5 * time.Second
	userInfoKey     = "userinfo"
)

type LoginController struct {
	BaseController
}

// Index 显示登录页面
func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

// DoLogin 处理登录请求
func (con LoginController) DoLogin(c *gin.Context) {
	captchaID := c.PostForm("captchaId")
	username := c.PostForm("username")
	password := c.PostForm("password")
	verifyValue := c.PostForm("verifyValue")

	if !models.VerifyCaptcha(captchaID, verifyValue) {
		con.Error(c, "验证码验证失败", "/admin/login")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), timeoutDuration)
	defer cancel()

	rbacClient := pbRbac.NewRbacLoginService("rbac", models.RbacClient)
	res, err := rbacClient.Login(ctx, &pbRbac.LoginRequest{
		Username: username,
		Password: models.Md5(password),
	})

	if err != nil {
		con.Error(c, fmt.Sprintf("登录失败: %v", err), "/admin/login")
		return
	}

	if !res.IsLogin {
		con.Error(c, "用户名或者密码错误", "/admin/login")
		return
	}

	if err := con.setUserSession(c, res.Userlist); err != nil {
		con.Error(c, fmt.Sprintf("设置会话失败: %v", err), "/admin/login")
		return
	}

	con.Success(c, "登录成功", "/admin")
}

// Captcha 生成验证码
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha(34, 100, 2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成验证码失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

// LoginOut 处理登出请求
func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userInfoKey)
	if err := session.Save(); err != nil {
		con.Error(c, fmt.Sprintf("退出登录失败: %v", err), "/admin")
		return
	}
	con.Success(c, "退出登录成功", "/admin/login")
}

// setUserSession 设置用户会话
func (con LoginController) setUserSession(c *gin.Context, userList interface{}) error {
	session := sessions.Default(c)
	userInfoSlice, err := json.Marshal(userList)
	if err != nil {
		return fmt.Errorf("序列化用户信息失败: %w", err)
	}
	session.Set(userInfoKey, string(userInfoSlice))
	return session.Save()
}
