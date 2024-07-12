package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loyalsfc/investrite/database"
	"github.com/loyalsfc/investrite/routes"
)

type Mide struct {
	Name string
}

func main() {
	//Initiate
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	db, err := database.InitDB()

	if err != nil {
		panic("failed to connect to database")
	}

	router := routes.InitRoutes(db)

	router.Run()
}
