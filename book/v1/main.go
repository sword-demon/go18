// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey;column:id"`
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200)" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

func (b *Book) TableName() string {
	return "books"
}

func setupDatabase() *gorm.DB {
	dsn := "root:admin888@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 开启 debug 模式
	db = db.Debug()
	db.AutoMigrate(&Book{})

	return db
}

type BookSet struct {
	Totoal int64   `json:"total"`
	List   []*Book `json:"list"`
}

type BookApiHandler struct{}

func NewBookApiHandler() *BookApiHandler {
	return &BookApiHandler{}
}

var db = setupDatabase()

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

	query := db

	kws := ctx.Query("keyword")
	if kws != "" {
		query = query.Where("title like ?", "%"+kws+"%")
	}

	err = query.Model(Book{}).Count(&set.Totoal).Offset(int((pn - 1) * ps)).Limit(int(ps)).Find(&set.List).Error
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

	book := &Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := db.Model(&Book{}).Create(&book).Error; err != nil {
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

	book := &Book{
		Title:  rq.Title,
		Author: rq.Author,
		Price:  rq.Price,
		IsSale: rq.IsSale,
	}
	if err := db.Model(&Book{}).Where("id = ?", rq.ID).Updates(&book).Error; err != nil {
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

	if err := db.Where("id = ?", rq.ID).Delete(&Book{}).Error; err != nil {
		ctx.JSON(500, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
}

func main() {
	r := gin.Default()

	h := NewBookApiHandler()
	r.GET("/books", h.list)
	r.POST("/books", h.add)
	r.PUT("/books", h.update)
	r.DELETE("/books", h.delete)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
