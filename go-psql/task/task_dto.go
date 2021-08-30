package task

type Task struct {
	Id              string `json:"id" db:"id"`
	WorkspaceUserId string `json:"workspace_user_id" db:"workspace_user_id"`
	WorkspaceId     string `json:"workspace_id" db:"workspace_id"`
	Title           string `json:"title" db:"title" validate:"required"`
}
