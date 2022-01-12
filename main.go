package main

import (
	"blog_gin/database"
	"blog_gin/routers"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()
	// 静态资源
	router.Static("/static", "./static")
	router.Run(":8081")
}
