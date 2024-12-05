package product

import (
	"github.com/kataras/iris/v12"
)

func createProduct(ctx iris.Context) {
	var product Product

	if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	if err := db.Create(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(product)
}

func getAllProducts(ctx iris.Context) {
	var products []Product
	db.Preload("Brand").Preload("Tags").Find(&products)
	ctx.JSON(products)
}

func UpdateProduct(ctx iris.Context) {
	var product Product
	productID, _ := ctx.Params().GetInt("id")

	if err := db.First(&product, productID).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Product not found"})
		return
	}

	if err := ctx.ReadJSON(&product); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	if err := db.Save(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Error updating product"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(product)
}

func DeleteProduct(ctx iris.Context) {
	var product Product
	productID, _ := ctx.Params().GetInt("id")

	if err := db.First(&product, productID).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Product not found"})
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Error updating product"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Product deleted successfully"})
}
