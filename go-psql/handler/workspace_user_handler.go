package handler

import (
	"fmt"
	"go-psql/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkspaceUserHandler struct {
	workspaceUserService service.WorkspaceUserService
}

//factory

func NewWorkspaceUserHandler(workspaceUserService service.WorkspaceUserService) WorkspaceUserHandler {
	return WorkspaceUserHandler{workspaceUserService: workspaceUserService}
}

//methods

func (wuh *WorkspaceUserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("uga booga?")
	user := wuh.workspaceUserService.GetUser(id)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "shit broke???"})
		return
	}
	c.JSON(http.StatusOK, user)
}
