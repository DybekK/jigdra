package workspace

type Workspace struct {
	Id              string `json:"id" db:"id"`
	WorkspaceUserId string `json:"workspaceUserId" db:"workspace_user_id"`
}
