package main

import (
	"authwithtoken/domain/auth/repo"
	"authwithtoken/domain/auth/usecase"
	appInit "authwithtoken/init"
	"fmt"

	handler "authwithtoken/domain/auth/handler/http"

	_ "authwithtoken/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	appInit.StartAppInit()
}

func main() {
	db, err := appInit.ConnectToPgServer()

	if err != nil {
		fmt.Println("error connect to database")
	}

	defer db.Close()

	//router
	e := echo.New()

	//middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	//depedency
	authRepo := repo.NewAuthRepo(db)

	authUc := usecase.NewAuthUsecase(authRepo)

	handler.NewAuthHander(e, authUc)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
