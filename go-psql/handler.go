package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("uga booga?")
	user := workspaceUserService.FindByUserId(id)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "shit broke???"})
		return
	}
	c.JSON(http.StatusOK, user)
}
