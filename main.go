package main

import (
	"repopatterngin/controller"
	"repopatterngin/helper"
	"repopatterngin/model"
	"repopatterngin/repository"
	"repopatterngin/router"
)

var db = helper.GetConnection()

func main() {

	db.AutoMigrate(&model.Product{})
	productRespository := repository.StartProductRepository()
	productcontroller := controller.StartProductController(productRespository)

	router.Run(productcontroller)
	
}
