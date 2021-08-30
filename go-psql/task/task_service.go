package task

import (
	"github.com/google/uuid"
)

type TaskService struct {
	taskRepository TaskRepository
}

//factory
func NewTaskService(taskRepository TaskRepository) TaskService {
	return TaskService{taskRepository: taskRepository}
}

//methods
func (ts *TaskService) CreateTask(workspace_user_id, workspace_id, title string) (*string, error) {
	var task Task
	task.Id = uuid.NewString()
	task.WorkspaceUserId = workspace_user_id
	task.WorkspaceId = workspace_id
	task.Title = title
	t, err := ts.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}
	return &t.Id, nil
}

func (ts *TaskService) GetTask(task_id string) (*Task, error) {
	return ts.taskRepository.ReadByTaskId(task_id)
}

func (ts *TaskService) GetUserTasks(workspace_user_id string) ([]Task, error) {
	return ts.taskRepository.ReadByUserId(workspace_user_id)
}
