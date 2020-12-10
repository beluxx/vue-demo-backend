package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()

	app.Listen(":8866")
}

type articlePlayload struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Status    bool   `json:"status"`
}

func newApp() *iris.Application {
	app := iris.Default()
	app.Get("/", index)

	api := app.Party("/api")
	{
		api.Get("/title", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"title": "Hello Vue~~~",
			})
		})
		api.Get("/languages", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"languages": []string{"PHP", "Javascript", "Golang"},
			})
		})
		api.Post("/post", func(ctx iris.Context) {
			var p articlePlayload

			err := ctx.ReadBody(&p)
			// var errs string
			if err != nil {
				ctx.JSON(iris.Map{
					"errors": err.Error(),
					"ok":     "",
				})
				return
			}
			fmt.Println("forms=", p)

			// for data := range forms {
			// 	fmt.Println(data)
			// 	// for field, value := range data {
			// 	// 	if string(value) == "" {
			// 	// 		errs[field] = []string{fmt.Sprintf("%s can't be empty", field)}
			// 	// 	}
			// 	// }
			// }
			ctx.JSON(iris.Map{
				"errors": "",
				"ok":     "success",
			})
		})
		api.Get("/post", func(ctx iris.Context) {
			var posts []interface{}
			for i := 0; i < 10; i++ {
				p := articlePlayload{
					Title:     "title" + strconv.Itoa(i),
					Author:    "author" + strconv.Itoa(i),
					Content:   "content" + strconv.Itoa(i),
					CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
					Status:    checkStatus(i),
				}
				pj, err := json.Marshal(p)
				var post map[string]interface{}
				err = json.Unmarshal(pj, &post)
				// fmt.Println("原始的", p)
				// fmt.Println("json化后", post)
				if err != nil {
					ctx.JSON(iris.Map{
						"errors": "",
						"ok":     "success",
						"posts":  posts,
					})
					return
				} else {
					posts = append(posts, post)
				}

			}

			ctx.JSON(iris.Map{
				"errors": "",
				"ok":     "success",
				"posts":  posts,
			})
		})

	}

	return app
}

func checkStatus(i int) (status bool) {
	if i%2 == 0 {
		status = true
	} else {
		status = false
	}
	return
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Index Page</h1>")
}
