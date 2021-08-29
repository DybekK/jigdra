package handler

import (
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
	user := wuh.workspaceUserService.GetUser(id)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
