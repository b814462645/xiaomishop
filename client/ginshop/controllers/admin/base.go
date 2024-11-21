package admin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BaseController 提供基本的控制器功能
type BaseController struct{}

// Success 渲染成功页面
func (con BaseController) Success(c *gin.Context, message string, redirectUrl string) {
	con.renderResult(c, "success", message, redirectUrl)
}

// Error 渲染错误页面
func (con BaseController) Error(c *gin.Context, message string, redirectUrl string) {
	con.renderResult(c, "error", message, redirectUrl)
}

// renderResult 是一个通用的渲染方法，用于渲染结果页面
func (con BaseController) renderResult(c *gin.Context, resultType, message, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/"+resultType+".html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
		"timestamp":   time.Now().Unix(),
	})
}

// JSON 返回JSON格式的响应
func (con BaseController) JSON(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}

// Redirect 执行重定向
func (con BaseController) Redirect(c *gin.Context, code int, location string) {
	c.Redirect(code, location)
}
