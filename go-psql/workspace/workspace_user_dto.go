package workspace

import "github.com/google/uuid"

type WorkspaceUser struct {
	Id       uuid.UUID `json:"id" db:"id"`
	UserId   string    `json:"userId" db:"user_id"`
	Nickname string    `json:"nickname" db:"nickname"`
}
