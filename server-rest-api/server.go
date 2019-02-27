package main

// Copyrights All right reserved RÉMY DEME

import (
	"server-project/server-rest-api/server-utilities/handler"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	//    "server-project/server-utilities/database"
	"server-project/server-rest-api/server-utilities/service"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

// GO Webserver

func main() {

	app := iris.New()
	//    defer database.DB.Close()
	//Middleware for JWt
	var jwtHandler = jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return service.SignKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
	//app.Use(jwtHandler.Serve)
	// error handling
	app.OnErrorCode(iris.StatusNotFound, handler.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, handler.InternalServerError)

	// Method Get
	// info: Use to initialise a session and get a token
	app.Get("/set", jwtHandler.Serve, handler.Set)
	app.Post("/login", handler.Login)
	app.Get("/logout", handler.Logout)
	app.Get("/delete", handler.Delete)
	app.Run(iris.Addr(":3000"))
}
