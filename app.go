package main

import (
	"github.com/rhperera/marvel-comic-api/config"
	_ "github.com/rhperera/marvel-comic-api/docs"
	"github.com/rhperera/marvel-comic-api/server"
)

// @title MarvelComicsAPI
// @description Retrieves information on Marvel characters
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main()  {
	config.Init()
	server.Init()
	server.InitAPI()
	server.Connect("8080")
}


