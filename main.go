package main

import (
	"Gin_book/database"
	"Gin_book/router"
)

func main() {
	//
	database.InitMysql()
	router := router.InitRouter()
	//静态资源
	router.Run(":8085")
}
