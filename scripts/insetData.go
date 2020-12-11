package scripts

import (
	"vue_demo_backend/models"
)

// InsertPost 插入一条文章数据
func InsertPost() error {
	post := models.Post{
		Title:   "Hello Golang",
		Author:  "luzp",
		Content: `Golang is a very good code language. \n Everybody should to learn it.^_^`,
		Status:  true,
	}

	db := models.DB()

	if err := db.FirstOrCreate(&post).Error; err != nil {
		// log.Printf("created one record failed: %s", err.Error)
		return err
	}
	return nil
}
