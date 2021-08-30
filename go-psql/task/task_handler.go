package task

import (
	"go-psql/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService    TaskService
	authMiddleware middleware.AuthMiddleware
}

//factory
func NewTaskHandler(taskService TaskService, authMiddleware middleware.AuthMiddleware) TaskHandler {
	return TaskHandler{
		taskService:    taskService,
		authMiddleware: authMiddleware,
	}
}

//methods
func (th *TaskHandler) CreateTask(c *gin.Context) {
	user, err := th.authMiddleware.GetWorkspaceUserFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	task_id, err := th.taskService.CreateTask(user.Id, user.WorkspaceId, task.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"task_id": task_id})

}

func (th *TaskHandler) GetTask(c *gin.Context) {
	task_id := c.Param("id")
	task, err := th.taskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (th *TaskHandler) GetUserTasks(c *gin.Context) {
	workspace_user_id := c.Param("uuid")
	tasks, err := th.taskService.GetUserTasks(workspace_user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}
