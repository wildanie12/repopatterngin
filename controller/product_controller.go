package controller

import (
	"context"
	"net/http"
	"repopatterngin/model"
	"repopatterngin/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductsController struct {
	Product repository.ProductRepository
}

func StartProductController(productcontroller repository.ProductRepository) ProductControler {
	return &ProductsController{Product: productcontroller}
}

func (controller *ProductsController) Create(c *gin.Context) {
	product := model.Product{}
	ctx := context.Background()

	c.Bind(&product)

	prod, err := controller.Product.Insert(ctx, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": prod,
	})
}

func (controller *ProductsController) Update(c *gin.Context) {
	product := model.Product{}
	productUpdate := model.ProductUpdate{}
	ctx := context.Background()
	GetId := c.Param("id")
	id, _ := strconv.Atoi(GetId)

	prod, err := controller.Product.Update(ctx, id, product, productUpdate)
	if err != nil {
		panic(err)
	}

	errOr := c.BindJSON(&prod)
	if errOr != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": prod,
	})
}

func (controller *ProductsController) FindAll(c *gin.Context) {
	ctx := context.Background()

	prod := controller.Product.FindAll(ctx)

	c.JSON(http.StatusOK, gin.H{
		"data": prod,
	})

}

func (controller *ProductsController) FindById(c *gin.Context) {
	product := model.Product{}
	ctx := context.Background()

	GetId := c.Param("id")
	id, _ := strconv.Atoi(GetId)

	prod, err := controller.Product.FindId(ctx, id, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": prod,
	})
}

func (controller *ProductsController) Delete(c *gin.Context) {
	ctx := context.Background()
	GetId := c.Param("id")
	id, _ := strconv.Atoi(GetId)

	prod, err := controller.Product.Delete(ctx, id)

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": prod,
	})
}
