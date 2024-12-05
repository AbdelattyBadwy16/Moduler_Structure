package brand

import "github.com/kataras/iris/v12"

func createBrand(ctx iris.Context) {
	var brand Brand

	if err := ctx.ReadJSON(&brand); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	if err := db.Create(&brand).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(brand)
}

func getALlBrands(ctx iris.Context) {
	var brands []Brand
	db.Find(&brands)
	ctx.JSON(brands)
}
