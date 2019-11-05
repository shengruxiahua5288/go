package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"userMgmtDemo/controller"
	"userMgmtDemo/db"
)

func main() {
	new(db.BaseGorm).InitDB()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)


	//使用go-json-rest自带的中间件IfMiddleware筛选判断


	router, err := rest.MakeRouter(
        rest.Post("/users",new (controller.UserController).InsertUser),
		rest.Get("/users/:id", new (controller.UserController).QueryUserById),
		rest.Get("/users",new (controller.UserController).QueryAllUser),
		rest.Delete("/users/:id", new (controller.UserController).DeleteUser),
		rest.Put("/users",new (controller.UserController).UpdateUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
