package admin

import (
	"ginshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	var focusList []models.Focus
	if err := models.DB.Find(&focusList).Error; err != nil {
		con.Error(c, "获取轮播图列表失败", "/admin")
		return
	}
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {
	focus := models.Focus{
		Title:   c.PostForm("title"),
		Link:    c.PostForm("link"),
		AddTime: int(models.GetUnix()),
	}

	var err error
	if focus.FocusType, err = models.Int(c.PostForm("focus_type")); err != nil {
		con.Error(c, "非法的焦点类型", "/admin/focus/add")
		return
	}
	if focus.Sort, err = models.Int(c.PostForm("sort")); err != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
		return
	}
	if focus.Status, err = models.Int(c.PostForm("status")); err != nil {
		con.Error(c, "非法的状态值", "/admin/focus/add")
		return
	}

	if focus.FocusImg, err = models.UploadImg(c, "focus_img"); err != nil {
		con.Error(c, "上传图片失败", "/admin/focus/add")
		return
	}

	if err := models.DB.Create(&focus).Error; err != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}
	var focus models.Focus
	if err := models.DB.First(&focus, id).Error; err != nil {
		con.Error(c, "轮播图不存在", "/admin/focus")
		return
	}
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "非法请求", "/admin/focus")
		return
	}

	var focus models.Focus
	if err := models.DB.First(&focus, id).Error; err != nil {
		con.Error(c, "轮播图不存在", "/admin/focus")
		return
	}

	focus.Title = c.PostForm("title")
	focus.Link = c.PostForm("link")

	if focus.FocusType, err = models.Int(c.PostForm("focus_type")); err != nil {
		con.Error(c, "非法的焦点类型", "/admin/focus/edit?id="+models.String(id))
		return
	}
	if focus.Sort, err = models.Int(c.PostForm("sort")); err != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
		return
	}
	if focus.Status, err = models.Int(c.PostForm("status")); err != nil {
		con.Error(c, "非法的状态值", "/admin/focus/edit?id="+models.String(id))
		return
	}

	if focusImg, err := models.UploadImg(c, "focus_img"); err == nil && focusImg != "" {
		focus.FocusImg = focusImg
	}

	if err := models.DB.Save(&focus).Error; err != nil {
		con.Error(c, "修改数据失败", "/admin/focus/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改轮播图成功", "/admin/focus")
	}
}

func (con FocusController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
		return
	}

	if err := models.DB.Delete(&models.Focus{}, id).Error; err != nil {
		con.Error(c, "删除数据失败", "/admin/focus")
	} else {
		con.Success(c, "删除数据成功", "/admin/focus")
	}
}
