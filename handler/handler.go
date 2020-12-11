package handler

import (
	"vue_demo_backend/models"
	"vue_demo_backend/scripts"

	"github.com/kataras/iris/v12"
)

var db = models.DB()

// GetPosts 获取所有的文章
func GetPosts(ctx iris.Context) {
	// db := models.DB()
	var posts []models.Post

	db.Find(&posts)
	ctx.Application().Logger().Println(posts)
	// ctx.Application().Logger().Println(result)

	// var posts []interface{}
	// for a := range ars {
	// 	p, err := json.Marshal(a)
	// 	var post map[string]interface{}
	// 	err = json.Unmarshal(p, &post)
	// 	if err != nil {
	// 		ctx.JSON(iris.Map{
	// 			"code":  iris.StatusBadRequest,
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	posts = append(posts, post)
	// }

	ctx.JSON(iris.Map{
		"code":  iris.StatusOK,
		"posts": posts,
	})
}

// GetPost 获取指定id的文章
func GetPost(ctx iris.Context) {
	var article models.Post
	id, _ := ctx.Params().GetUint("id")
	err := db.Where("id = ?", int(id)).First(&article).Error
	if err != nil {
		ctx.JSON(iris.Map{
			"code":  iris.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"data": article,
	})
}

// InsertPost 新增文章
func InsertPost(ctx iris.Context) {
	var article models.Post
	err := ctx.ReadBody(&article)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":  iris.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	ctx.Application().Logger().Println(article)
	ctx.JSON(iris.Map{
		"code": iris.StatusOK,
	})
}

// InsertTmpPost 临时插入一条post数据
func InsertTmpPost(ctx iris.Context) {
	if err := scripts.InsertPost(); err != nil {
		ctx.JSON(iris.Map{
			"code":  iris.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(iris.Map{
		"code": iris.StatusOK,
	})
}

// DeletePost 删除文章
func DeletePost(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	if id == 0 {
		ctx.JSON(iris.Map{
			"code":  iris.StatusBadRequest,
			"error": "要删除的文章其id不能为空且不能为0",
		})
		return
	}
	var post models.Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		ctx.JSON(iris.Map{
			"code":  iris.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	db.Delete(&post)
	ctx.JSON(iris.Map{
		"code": iris.StatusOK,
		"data": post,
	})
}
