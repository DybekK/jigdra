package repository

import (
	"context"
	"go-psql/database"
	"go-psql/dto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type WorkspaceUserRepository struct {
	posgresDatabase *database.PostgresDatabase
}

//factory

func NewWorkspaceUserRepo(posgresDatabase database.PostgresDatabase) WorkspaceUserRepository {
	return WorkspaceUserRepository{posgresDatabase: &posgresDatabase}
}

//methods

func (wur *WorkspaceUserRepository) Create(user dto.WorkspaceUser) error {
	tx, err := wur.posgresDatabase.Connection.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	_, err = tx.Exec(context.Background(), `INSERT INTO workspaceusers (id, user_id, nickname) 
											VALUES ($1, $2, $3)`, user.Id, user.UserId, user.Nickname)
	if err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	return err
}

func (wur *WorkspaceUserRepository) Read(id string) *dto.WorkspaceUser {
	var uuid uuid.UUID
	var userid string
	var nickname string
	err := wur.posgresDatabase.Connection.QueryRow(context.Background(), `SELECT * FROM workspaceusers 
														WHERE user_id=$1`, id).Scan(&uuid, &userid, &nickname)
	if err == pgx.ErrNoRows {
		return nil
	}
	// var uuid uuid.UUID
	// row.Scan(&uuid)
	// user.Id = uuid
	user := dto.WorkspaceUser{
		Id:       uuid,
		UserId:   userid,
		Nickname: nickname,
	}
	return &user
}
