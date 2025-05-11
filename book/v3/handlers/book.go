package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/go18/book/v3/config"
	"github.com/sword-demon/go18/book/v3/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type (
	Base struct {
		Title  string  `json:"title" binding:"required"`
		Author string  `json:"author"  binding:"required"`
		Price  float64 `json:"price"  binding:"required"`
		IsSale *bool   `json:"is_sale"`
	}
	CreateBookReq struct {
		Base
	}

	UpdateBookReq struct {
		ID uint `json:"id" binding:"required"`
		Base
	}
	DeleteBookReq struct {
		ID uint `json:"id" binding:"required"`
	}
)

type BookSet struct {
	Totoal int64          `json:"total"`
	List   []*models.Book `json:"list"`
}

type BookApiHandler struct{}

func (h *BookApiHandler) Registry(r gin.IRouter) {
	r.GET("/books", h.list)
	r.POST("/books", h.add)
	r.PUT("/books", h.update)
	r.DELETE("/books", h.delete)
}

func NewBookApiHandler() *BookApiHandler {
	return &BookApiHandler{}
}

func (h *BookApiHandler) list(ctx *gin.Context) {
	set := &BookSet{}
	pageSize := ctx.Query("page_size")
	page := ctx.Query("page")

	pn, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	ps, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
	}

	query := config.DB()

	kws := ctx.Query("keyword")
	if kws != "" {
		query = query.Where("title like ?", "%"+kws+"%")
	}

	err = query.Model(models.Book{}).Count(&set.Totoal).Offset(int((pn - 1) * ps)).Limit(int(ps)).Find(&set.List).Error
	if err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(200, gin.H{"code": 500, "msg": "暂无数据"})
		return
	}

	ctx.JSON(200, set)
}

func (h *BookApiHandler) add(ctx *gin.Context) {
	var rq CreateBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	book := &models.Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := config.DB().Model(&models.Book{}).Create(&book).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "插入成功", "data": book.ID})
}

func (h *BookApiHandler) update(ctx *gin.Context) {
	var rq UpdateBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	book := &models.Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := config.DB().Model(&models.Book{}).Where("id = ?", rq.ID).Updates(&book).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改", "data": book})
}

func (h *BookApiHandler) delete(ctx *gin.Context) {
	var rq DeleteBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	if err := config.DB().Where("id = ?", rq.ID).Delete(&models.Book{}).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
}
