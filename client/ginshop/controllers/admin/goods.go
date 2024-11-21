package admin

import (
	"fmt"
	"ginshop/models"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

var goodsForm struct {
	Title         string   `form:"title" binding:"required"`
	SubTitle      string   `form:"sub_title"`
	GoodsSn       string   `form:"goods_sn"`
	CateId        int      `form:"cate_id" binding:"required"`
	GoodsNumber   int      `form:"goods_number" binding:"required,min=0"`
	MarketPrice   float64  `form:"market_price" binding:"required,min=0"`
	Price         float64  `form:"price" binding:"required,min=0"`
	RelationGoods string   `form:"relation_goods"`
	GoodsAttr     string   `form:"goods_attr"`
	GoodsVersion  string   `form:"goods_version"`
	GoodsGift     string   `form:"goods_gift"`
	GoodsFitting  string   `form:"goods_fitting"`
	GoodsColor    []string `form:"goods_color"`
	GoodsKeywords string   `form:"goods_keywords"`
	GoodsDesc     string   `form:"goods_desc"`
	GoodsContent  string   `form:"goods_content"`
	IsDelete      int      `form:"is_delete"`
	IsHot         int      `form:"is_hot"`
	IsBest        int      `form:"is_best"`
	IsNew         int      `form:"is_new"`
	GoodsTypeId   int      `form:"goods_type_id"`
	Sort          int      `form:"sort"`
	Status        int      `form:"status"`
}

func (con GoodsController) Index(c *gin.Context) {
	// 当前页数
	page, err := models.Int(c.Query("page"))
	if page == 0 || err != nil {
		page = 1
	}

	// 条件
	where := "is_delete = ?"
	args := []interface{}{0}

	// 获取keyword
	keyword := c.Query("keyword")
	if len(keyword) > 0 {
		where += " AND title LIKE ?"
		args = append(args, "%"+keyword+"%")
	}

	// 每页查询的数量
	pageSize := 8

	goodsList := []models.Goods{}
	models.DB.Where(where, args...).Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&goodsList)

	// 获取总数量
	var count int64
	models.DB.Where(where, args...).Table("goods").Count(&count)

	// 判断最后一页有没有数据 如果没有跳转到第一页
	if len(goodsList) > 0 {
		c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
			"goodsList":  goodsList,
			"totalPages": math.Ceil(float64(count) / float64(pageSize)),
			"page":       page,
			"keyword":    keyword,
		})
	} else {
		if page != 1 {
			c.Redirect(302, "/admin/goods")
		} else {
			c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
				"goodsList":  goodsList,
				"totalPages": math.Ceil(float64(count) / float64(pageSize)),
				"page":       page,
				"keyword":    keyword,
			})
		}
	}
}

func (con GoodsController) Add(c *gin.Context) {
	//获取商品分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	goodsColorList := []models.GoodsColor{}
	models.DB.Find(&goodsColorList)

	//获取商品规格包装
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := models.Int(c.Query("cateId"))
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err2 := models.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  goodsTypeAttributeList,
		})
	}
}

func (con GoodsController) DoAdd(c *gin.Context) {
	if err := c.ShouldBind(&goodsForm); err != nil {
		con.Error(c, "输入数据无效: "+err.Error(), "/admin/goods/add")
		return
	}

	// 2. 准备商品数据
	goods := models.Goods{
		Title:         goodsForm.Title,
		SubTitle:      goodsForm.SubTitle,
		GoodsSn:       goodsForm.GoodsSn,
		CateId:        goodsForm.CateId,
		ClickCount:    100,
		GoodsNumber:   goodsForm.GoodsNumber,
		MarketPrice:   goodsForm.MarketPrice,
		Price:         goodsForm.Price,
		RelationGoods: goodsForm.RelationGoods,
		GoodsAttr:     goodsForm.GoodsAttr,
		GoodsVersion:  goodsForm.GoodsVersion,
		GoodsGift:     goodsForm.GoodsGift,
		GoodsFitting:  goodsForm.GoodsFitting,
		GoodsKeywords: goodsForm.GoodsKeywords,
		GoodsDesc:     goodsForm.GoodsDesc,
		GoodsContent:  goodsForm.GoodsContent,
		IsDelete:      goodsForm.IsDelete,
		IsHot:         goodsForm.IsHot,
		IsBest:        goodsForm.IsBest,
		IsNew:         goodsForm.IsNew,
		GoodsTypeId:   goodsForm.GoodsTypeId,
		Sort:          goodsForm.Sort,
		Status:        goodsForm.Status,
		AddTime:       int(models.GetUnix()),
		GoodsColor:    strings.Join(goodsForm.GoodsColor, ","),
	}

	// 3. 处理图片上传
	if goodsImg, err := models.UploadImg(c, "goods_img"); err == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		if models.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				models.ResizeGoodsImage(goodsImg)
			}()
		}
	}

	// 4. 使用事务处理数据库操作
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// 4.1 添加商品
		if err := tx.Create(&goods).Error; err != nil {
			return err
		}

		// 4.2 处理图库信息
		goodsImageList := c.PostFormArray("goods_image_list")
		var goodsImages []models.GoodsImage
		for _, v := range goodsImageList {
			goodsImages = append(goodsImages, models.GoodsImage{
				GoodsId: goods.Id,
				ImgUrl:  v,
				Sort:    10,
				Status:  1,
				AddTime: int(models.GetUnix()),
			})
		}
		if err := tx.CreateInBatches(goodsImages, 100).Error; err != nil {
			return err
		}

		// 4.3 处理规格包装
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		var goodsAttrs []models.GoodsAttr
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, err := models.Int(attrIdList[i])
			if err != nil {
				continue
			}

			var goodsTypeAttribute models.GoodsTypeAttribute
			if err := tx.First(&goodsTypeAttribute, goodsTypeAttributeId).Error; err != nil {
				continue
			}

			goodsAttrs = append(goodsAttrs, models.GoodsAttr{
				GoodsId:         goods.Id,
				AttributeTitle:  goodsTypeAttribute.Title,
				AttributeType:   goodsTypeAttribute.AttrType,
				AttributeId:     goodsTypeAttribute.Id,
				AttributeCateId: goodsTypeAttribute.CateId,
				AttributeValue:  attrValueList[i],
				Status:          1,
				Sort:            10,
				AddTime:         int(models.GetUnix()),
			})
		}

		if err := tx.CreateInBatches(goodsAttrs, 100).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		con.Error(c, "增加商品失败: "+err.Error(), "/admin/goods/add")
		return
	}

	wg.Wait()

	con.Success(c, "增加商品成功", "/admin/goods")
}

// 修改
func (con GoodsController) Edit(c *gin.Context) {

	// 1、获取要修改的商品数据
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
	}
	goods := models.Goods{}
	if err := models.DB.First(&goods, id).Error; err != nil {
		con.Error(c, "商品不存在", "/admin/goods")
		return
	}

	// 2、获取商品分类
	var goodsCateList []models.GoodsCate
	if err := models.DB.Where("pid = ?", 0).Preload("GoodsCateItems").Find(&goodsCateList).Error; err != nil {
		con.Error(c, "获取商品分类失败", "/admin/goods")
		return
	}

	// 3、获取所有颜色 以及选中的颜色
	goodsColorMap := make(map[string]bool)
	for _, v := range strings.Split(goods.GoodsColor, ",") {
		goodsColorMap[v] = true
	}

	var goodsColorList []models.GoodsColor
	if err := models.DB.Find(&goodsColorList).Error; err != nil {
		con.Error(c, "获取颜色信息失败", "/admin/goods")
		return
	}

	for i := range goodsColorList {
		goodsColorList[i].Checked = goodsColorMap[models.String(goodsColorList[i].Id)]
	}

	// 4、商品的图库信息
	goodsImageList := []models.GoodsImage{}
	if err := models.DB.Where("goods_id=?", goods.Id).Find(&goodsImageList).Error; err != nil {
		con.Error(c, "获取图库信息失败", "/admin/goods")
		return
	}

	// 5、获取商品类型
	goodsTypeList := []models.GoodsType{}
	if err := models.DB.Find(&goodsTypeList).Error; err != nil {
		con.Error(c, "获取商品类型失败", "/admin/goods")
		return
	}

	// 6、获取规格信息
	goodsAttr := []models.GoodsAttr{}
	if err := models.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr).Error; err != nil {
		con.Error(c, "获取规格信息失败", "/admin/goods")
		return
	}
	goodsAttrStr := con.buildGoodsAttrStr(goodsAttr)

	//获取上一页的地址
	// fmt.Println(c.Request.Referer())

	c.HTML(http.StatusOK, "admin/goods/edit.html", gin.H{
		"goods":          goods,
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
		"goodsAttrStr":   goodsAttrStr,
		"goodsImageList": goodsImageList,
		"prevPage":       c.Request.Referer(), //获取上一页的地址
	})
}

// 构建商品属性字符串
func (con GoodsController) buildGoodsAttrStr(goodsAttr []models.GoodsAttr) string {
	var goodsAttrStr strings.Builder
	for _, v := range goodsAttr {
		goodsAttrStr.WriteString("<li><span>" + v.AttributeTitle + ": </span>")
		goodsAttrStr.WriteString("<input type=\"hidden\" name=\"attr_id_list\" value=\"" + strconv.Itoa(v.AttributeId) + "\" />")

		switch v.AttributeType {
		case 1:
			goodsAttrStr.WriteString("<input type=\"text\" name=\"attr_value_list\" value=\"" + v.AttributeValue + "\" />")
		case 2:
			goodsAttrStr.WriteString("<textarea cols=\"50\" rows=\"3\" name=\"attr_value_list\">" + v.AttributeValue + "</textarea>")
		default:
			goodsAttrStr.WriteString(con.buildSelectOptions(v))
		}

		goodsAttrStr.WriteString("</li>")
	}
	return goodsAttrStr.String()
}

// 构建选择选项
func (con GoodsController) buildSelectOptions(attr models.GoodsAttr) string {
	var goodsTypeAttribute models.GoodsTypeAttribute
	if err := models.DB.First(&goodsTypeAttribute, attr.AttributeId).Error; err != nil {
		return ""
	}

	var selectStr strings.Builder
	selectStr.WriteString("<select name=\"attr_value_list\">")

	attrValues := strings.Split(goodsTypeAttribute.AttrValue, "\n")
	for _, value := range attrValues {
		selected := ""
		if value == attr.AttributeValue {
			selected = " selected"
		}
		selectStr.WriteString("<option value=\"" + value + "\"" + selected + ">" + value + "</option>")
	}

	selectStr.WriteString("</select>")
	return selectStr.String()
}

func (con GoodsController) DoEdit(c *gin.Context) {

	//1、获取表单提交过来的数据
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/goods")
		return
	}
	//获取上一页的地址
	prevPage := c.PostForm("prevPage")

	if err := c.ShouldBind(&goodsForm); err != nil {
		con.Error(c, "表单数据无效", "/admin/goods/edit?id="+models.String(id))
		return
	}

	// 3. 更新商品信息
	goods := models.Goods{Id: id}
	if err := models.DB.First(&goods).Error; err != nil {
		con.Error(c, "商品不存在", "/admin/goods")
		return
	}

	goods.Title = goodsForm.Title
	goods.SubTitle = goodsForm.SubTitle
	goods.GoodsSn = goodsForm.GoodsSn
	goods.CateId = goodsForm.CateId
	goods.GoodsNumber = goodsForm.GoodsNumber
	goods.MarketPrice = goodsForm.MarketPrice
	goods.Price = goodsForm.Price
	goods.RelationGoods = goodsForm.RelationGoods
	goods.GoodsAttr = goodsForm.GoodsAttr
	goods.GoodsVersion = goodsForm.GoodsVersion
	goods.GoodsGift = goodsForm.GoodsGift
	goods.GoodsFitting = goodsForm.GoodsFitting
	goods.GoodsKeywords = goodsForm.GoodsKeywords
	goods.GoodsDesc = goodsForm.GoodsDesc
	goods.GoodsContent = goodsForm.GoodsContent
	goods.IsDelete = goodsForm.IsDelete
	goods.IsHot = goodsForm.IsHot
	goods.IsBest = goodsForm.IsBest
	goods.IsNew = goodsForm.IsNew
	goods.GoodsTypeId = goodsForm.GoodsTypeId
	goods.Sort = goodsForm.Sort
	goods.Status = goodsForm.Status
	goods.GoodsColor = strings.Join(goodsForm.GoodsColor, ",")

	//4、上传图片   生成缩略图
	if goodsImg, err := models.UploadImg(c, "goods_img"); err == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
		if models.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				models.ResizeGoodsImage(goodsImg)
			}()

		}
	}

	// 5. 保存商品信息
	if err := models.DB.Save(&goods).Error; err != nil {
		con.Error(c, "修改失败", "/admin/goods/edit?id="+models.String(id))
		return
	}

	//6、修改规格包装  1、删除当前商品下面的规格包装   2、重新执行增加
	var handleErr error
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := con.handleGoodsImageAndAttr(c, goods.Id); err != nil {
			con.Error(c, "处理商品图片和属性失败", "/admin/goods/edit?id="+models.String(id))
			handleErr = err
		}
	}()
	wg.Wait()

	if handleErr != nil {
		con.Error(c, "处理商品图片和属性失败", "/admin/goods/edit?id="+models.String(id))
		return
	}

	if len(prevPage) > 0 {
		con.Success(c, "修改数据成功", prevPage)
	} else {
		con.Success(c, "修改数据成功", "/admin/goods")
	}

}

func (con GoodsController) handleGoodsImageAndAttr(c *gin.Context, goodsId int) error {
	// 处理图库信息
	goodsImageList := c.PostFormArray("goods_image_list")
	for _, v := range goodsImageList {
		goodsImgObj := models.GoodsImage{
			GoodsId: goodsId,
			ImgUrl:  v,
			Sort:    10,
			Status:  1,
			AddTime: int(models.GetUnix()),
		}
		models.DB.Create(&goodsImgObj)
	}

	// 处理规格包装
	if err := models.DB.Where("goods_id = ?", goodsId).Delete(&models.GoodsAttr{}).Error; err != nil {
		return fmt.Errorf("删除旧的商品属性失败: %v", err)
	}

	attrIdList := c.PostFormArray("attr_id_list")
	attrValueList := c.PostFormArray("attr_value_list")
	for i := 0; i < len(attrIdList); i++ {
		goodsTypeAttributeId, err := models.Int(attrIdList[i])
		if err != nil {
			continue
		}

		var goodsTypeAttribute models.GoodsTypeAttribute
		if err := models.DB.First(&goodsTypeAttribute, goodsTypeAttributeId).Error; err != nil {
			continue
		}

		goodsAttr := models.GoodsAttr{
			GoodsId:         goodsId,
			AttributeTitle:  goodsTypeAttribute.Title,
			AttributeType:   goodsTypeAttribute.AttrType,
			AttributeId:     goodsTypeAttribute.Id,
			AttributeCateId: goodsTypeAttribute.CateId,
			AttributeValue:  attrValueList[i],
			Status:          1,
			Sort:            10,
			AddTime:         int(models.GetUnix()),
		}
		if err := models.DB.Create(&goodsAttr).Error; err != nil {
			return fmt.Errorf("创建商品属性失败: %v", err)
		}
	}

	return nil
}

// 富文本编辑器上传图片
func (con GoodsController) EditorImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := models.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		if models.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				models.ResizeGoodsImage(imgDir)
				wg.Done()
			}()
			c.JSON(http.StatusOK, gin.H{
				"link": "/" + imgDir,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"link": models.GetSettingFromColumn("OssDomain") + imgDir,
			})
		}

	}
}

// 图库上传图片
func (con GoodsController) GoodsImageUpload(c *gin.Context) {
	//上传图片
	imgDir, err := models.UploadImg(c, "file") //注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
	} else {
		if models.GetOssStatus() != 1 {
			wg.Add(1)
			go func() {
				models.ResizeGoodsImage(imgDir)
				wg.Done()
			}()

		}
		c.JSON(http.StatusOK, gin.H{
			"link": imgDir,
		})

	}
}

// 修改商品图库关联的颜色
func (con GoodsController) ChangeGoodsImageColor(c *gin.Context) {
	//获取图片id 获取颜色id
	goodsImageId, err1 := models.Int(c.Query("goods_image_id"))
	colorId, err2 := models.Int(c.Query("color_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	models.DB.Find(&goodsImage)
	goodsImage.ColorId = colorId
	err3 := models.DB.Save(&goodsImage).Error
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新失败",
			"success": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result":  "更新成功",
			"success": true,
		})
	}

}

// 删除图库
func (con GoodsController) RemoveGoodsImage(c *gin.Context) {
	//获取图片id
	goodsImageId, err1 := models.Int(c.Query("goods_image_id"))
	goodsImage := models.GoodsImage{Id: goodsImageId}
	err2 := models.DB.Delete(&goodsImage).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除失败",
			"success": false,
		})
	} else {
		//删除图片
		// os.Remove()
		c.JSON(http.StatusOK, gin.H{
			"result":  "删除成功",
			"success": true,
		})
	}

}

// 删除数据
func (con GoodsController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goods")
	} else {
		goods := models.Goods{Id: id}
		models.DB.Find(&goods)
		goods.IsDelete = 1
		goods.Status = 0
		models.DB.Save(&goods)
		//获取上一页
		prevPage := c.Request.Referer()
		if len(prevPage) > 0 {
			con.Success(c, "删除数据成功", prevPage)
		} else {
			con.Success(c, "删除数据成功", "/admin/goods")
		}

	}

}
