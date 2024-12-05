package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	InitDB()
	app := iris.New()

	app.Get("/products", getAllProducts)
	app.Post("/products", createProduct)
	app.Patch("/products/{id:int}", UpdateProduct)
	app.Delete("/products/{id:int}", DeleteProduct)
	app.Get("/brands", getALlBrands)
	app.Post("/brands", createBrand)
	app.Get("/tags", getAllTags)
	app.Post("/tags", createTag)

	app.Listen(":8080")
}
