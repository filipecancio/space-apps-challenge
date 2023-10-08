package controller

import (
	"fmt"
	"net/http"

	"github.com/filipecancio/space-apps-challenge/odin-community/src/api/model"
	"github.com/gin-gonic/gin"
)

func InsertReport(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-16")
	var newid int
	rep := new(model.Report)
	err := c.BindJSON(&rep)
	fmt.Println(err)
	newid = model.Newreportservice().Create(*rep)
	c.IndentedJSON(http.StatusOK, struct{ id int }{newid})
}

func ListReport(c *gin.Context) {
	userId := c.Param("user")
	c.IndentedJSON(http.StatusOK, model.Newreportservice().List(userId))
}

func ListFire(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Newfireservice().List())
}
