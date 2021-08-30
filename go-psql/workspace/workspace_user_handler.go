package workspace

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkspaceUserHandler struct {
	workspaceUserService WorkspaceUserService
}

//factory

func NewWorkspaceUserHandler(workspaceUserService WorkspaceUserService) WorkspaceUserHandler {
	return WorkspaceUserHandler{workspaceUserService: workspaceUserService}
}

//methods

func (wuh *WorkspaceUserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := wuh.workspaceUserService.GetUser(id)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"err": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
