package main

import (
	"fmt"
	"log"

	"github.com/filipecancio/space-apps-challenge/odin-community/src/api/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/v1/report", controller.InsertReport)
	router.GET("/v1/user/:user/report", controller.ListReport)
	router.GET("/v1/fire", controller.ListFire)

	fmt.Println("ODIN Loading...")
	log.Fatal(router.Run(":4444"))
}
