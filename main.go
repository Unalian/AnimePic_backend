
package main

import (
	"Pic/src"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)


func main() {
	src.ConnectDB()
	defer src.DisconnectDb()
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	api := app.Party("/api", crs).AllowMethods(iris.MethodOptions)

	api.Post("/login", src.Login)

	api.Post("/register", src.Register)

	api.Post("/UserInfo", src.GetUser)

	api.Get("/img", src.GetImg)

	api.Post("/file", src.GetFileList)

	api.Post("/crwal", src.)

	_ = app.Run(iris.Addr(":15551"))

}