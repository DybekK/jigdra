package workspace

type WorkspaceUser struct {
	Id       string `json:"id" db:"id"`
	UserId   string `json:"userId" db:"user_id"`
	Nickname string `json:"nickname" db:"nickname"`
}
