package repository

import (
	"context"
	"go-psql/model/dto"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type WorkspaceUserRepo interface {
	Create(dto.WorkspaceUser) error
	Read(string) *dto.WorkspaceUser
	// Update() error
	// Delete() error
}

func NewWorkspaceUserRepo(conn *pgx.Conn) WorkspaceUserRepo {
	return &database{Connection: conn}
}

func (d *database) Create(user dto.WorkspaceUser) error {
	tx, err := d.Connection.Begin(context.Background())
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

func (d *database) Read(id string) *dto.WorkspaceUser {
	var uuid uuid.UUID
	var userid string
	var nickname string
	err := d.Connection.QueryRow(context.Background(), `SELECT * FROM workspaceusers 
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
