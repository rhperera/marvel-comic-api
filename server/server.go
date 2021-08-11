package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rhperera/marvel-comic-api/common"
	"github.com/rhperera/marvel-comic-api/services"
	"github.com/swaggo/echo-swagger"
	"log"
	"net/http"
)

var EchoCon *echo.Echo
var EchoRG *echo.Group

func Init() {
	EchoCon = echo.New()
	http.Handle("/", EchoCon)
}

/*
Calling this method will block the main thread
until further callbacks from the echo framework
*/
func Connect(port string) {
	if EchoCon == nil {
		log.Fatal(common.ErrorEchoServerInit)
		return
	}
	EchoCon.Logger.Fatal(EchoCon.Start(":" + port))
}

func InitAPI() {
	if EchoCon == nil {
		log.Fatal(common.ErrorEchoServerInit)
	}
	cacheService := &services.RedisCacheService{}
	cacheService.Connect()
	handler := NewHandler(cacheService, &services.MarvelCharacterAPI{})

	EchoRG = EchoCon.Group("/api/v1")

	EchoRG.GET("/swagger/*", echoSwagger.WrapHandler)
	EchoRG.GET("/characters", handler.GetAllCharacters)
	EchoRG.GET("/characters/:id", handler.GetCharacterById)
}

func defaultGetOk(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
