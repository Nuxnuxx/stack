package main

import (
	"stack/db"
	"stack/handler"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//FIXME: In production, the secret key of the CookieStore
// and database name would be obtained from a .env file

const (
	SECRET_KEY string = "secret"
	DB_NAME    string = "todo.db"
)

func main() {
	e := echo.New()

	// serve css and js
	e.Static("/", "static")

	// Ca log c'est tout
	e.Use(middleware.Logger())
	// what is this shit idk
	e.Use(middleware.Recover())
	// sessions but idk how it works
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	store, err := db.NewStore(DB_NAME)
	if err != nil {
		e.Logger.Fatal("failed to create store", err)
	}

	helloHandler := handler.HelloHandler{}

	e.GET("/", helloHandler.Show)

	e.Logger.Fatal(e.Start(":8080"))
}
