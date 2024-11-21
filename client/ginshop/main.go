package main

import (
	"ginshop/models"
	"ginshop/routers"
	"html/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 路由引擎
	r := gin.Default()

	// 自定义模板函数
	// 注意: 必须在加载模板之前设置自定义函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime":   models.UnixToTime,   // 将 Unix 时间戳转换为时间
		"UnixToTime64": models.UnixToTime64, // 将 64 位 Unix 时间戳转换为时间
		"Str2Html":     models.Str2Html,     // 将字符串转换为 HTML
		"FormatImg":    models.FormatImg,    // 格式化图片路径
		"Sub":          models.Sub,          // 减法运算
		"Mul":          models.Mul,          // 乘法运算
		"Substr":       models.Substr,       // 截取字符串
		"FormatAttr":   models.FormatAttr,   // 格式化属性
	})

	// 加载 HTML 模板
	// 注意: 必须在配置路由之前加载模板
	r.LoadHTMLGlob("templates/**/**/*")

	// 配置静态文件服务
	// 第一个参数是 URL 路径，第二个参数是文件系统路径
	r.Static("/static", "./static")

	// 创建基于 cookie 的会话存储
	// "secret111" 是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))

	// 配置会话中间件
	// 使用前面创建的 cookie 存储引擎
	// 注意: 可以根据需要替换为其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	// 初始化管理员路由
	routers.AdminRoutersInit(r)

	// 初始化 API 路由
	routers.ApiRoutersInit(r)

	// 初始化默认路由
	routers.DefaultRoutersInit(r)

	// 启动服务器
	// 默认监听 :8080 端口，可以通过环境变量或命令行参数修改
	r.Run()
}
