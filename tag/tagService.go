package tag

import "github.com/kataras/iris/v12"

func createTag(ctx iris.Context) {
	var tag Tag

	if err := ctx.ReadJSON(&tag); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	if err := db.Create(&tag).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(tag)
}

func getAllTags(ctx iris.Context) {
	var tags []Tag
	db.Find(&tags)
	ctx.JSON(tags)
}
