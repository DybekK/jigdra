package workspace

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WorkspaceFacadeHandler struct {
	workspaceFacade      WorkspaceFacade
	workspaceUserService WorkspaceUserService
}

//factory

func NewWorkspaceFacadeHandler(workspaceFacade WorkspaceFacade, workspaceUserService WorkspaceUserService) WorkspaceFacadeHandler {
	return WorkspaceFacadeHandler{workspaceFacade: workspaceFacade, workspaceUserService: workspaceUserService}
}

//methods

func (w *WorkspaceFacadeHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := w.workspaceUserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (w *WorkspaceFacadeHandler) CreateUserAndWorkspace(c *gin.Context) {
	type Body struct {
		UserId   string `json:"userId"`
		Nickname string `json:"nickname"`
	}
	var body Body
	c.BindJSON(&body)

	_, _, err := w.workspaceFacade.CreateUserAndWorkspace(body.UserId, body.Nickname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "workspace and user have been created"})
}
