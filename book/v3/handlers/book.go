package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sword-demon/go18/book/v3/config"
	"github.com/sword-demon/go18/book/v3/controllers"
	"github.com/sword-demon/go18/book/v3/models"
	"github.com/sword-demon/go18/book/v3/response"
	"gorm.io/gorm"
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
	GetBookRequest struct {
		BookNumber string `json:"book_number"`
	}
)

type BookSet struct {
	Total int64          `json:"total"`
	List  []*models.Book `json:"list"`
}

type BookApiHandler struct{}

func (h *BookApiHandler) Registry(r gin.IRouter) {
	r.GET("/books", h.list)
	r.POST("/books", h.add)
	r.PUT("/books", h.update)
	r.DELETE("/books", h.delete)
	r.GET("/books/:bn", h.getBook)
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
		response.Failed(ctx, err)
		return
	}
	ps, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	query := config.DB()

	kws := ctx.Query("keyword")
	if kws != "" {
		query = query.Where("title like ?", "%"+kws+"%")
	}

	err = query.Model(models.Book{}).Count(&set.Total).Offset(int((pn - 1) * ps)).Limit(int(ps)).Find(&set.List).Error
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OK(ctx, nil)
		return
	}

	response.OK(ctx, set)
}

func (h *BookApiHandler) add(ctx *gin.Context) {
	var rq CreateBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		response.Failed(ctx, err)
		return
	}

	book := &models.Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := config.DB().Model(&models.Book{}).Create(&book).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, book)
}

func (h *BookApiHandler) update(ctx *gin.Context) {
	var rq UpdateBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		response.Failed(ctx, err)
		return
	}

	book := &models.Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := config.DB().Model(&models.Book{}).Where("id = ?", rq.ID).Updates(&book).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, book)
}

func (h *BookApiHandler) delete(ctx *gin.Context) {
	var rq DeleteBookReq
	if err := ctx.ShouldBindJSON(&rq); err != nil {
		response.Failed(ctx, err)
		return
	}

	if err := config.DB().Where("id = ?", rq.ID).Delete(&models.Book{}).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, nil)
}

func (h *BookApiHandler) getBook(ctx *gin.Context) {
	bookC := controllers.NewBookController()
	req := &controllers.GetBookRequest{BookNumber: ctx.Param("bn")}
	book, err := bookC.GetBook(ctx, req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.OK(ctx, book)
}
