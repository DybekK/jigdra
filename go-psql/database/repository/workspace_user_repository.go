package repository

import (
	"context"
	"fmt"
	"go-psql/dto"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WorkspaceUserRepository struct {
	posgresDatabase *pgxpool.Pool
}

//factory

func NewWorkspaceUserRepo(posgresDatabase *pgxpool.Pool) WorkspaceUserRepository {
	return WorkspaceUserRepository{posgresDatabase: posgresDatabase}
}

//methods

func (wur *WorkspaceUserRepository) Create(user dto.WorkspaceUser) error {
	tx, err := wur.posgresDatabase.Begin(context.Background())
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
	var user dto.WorkspaceUser
	fmt.Println(id)
	row, err := wur.posgresDatabase.Query(context.Background(), `SELECT * FROM workspaceusers WHERE user_id=$1`, id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = pgxscan.ScanOne(&user, row)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &user
}
