package router

import (
	"repopatterngin/controller"

	"github.com/gin-gonic/gin"
)

func Run(routermy controller.ProductControler) {
	server := gin.Default()
	router := server.Group("/api/product")
	router.GET("", routermy.FindAll)
	router.GET("/:id", routermy.FindById)
	router.PUT("/:id", routermy.Update)
	router.POST("", routermy.Create)
	router.DELETE("/:id", routermy.Delete)

	server.Run()
}
