package workspace

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WorkspaceUserRepository struct {
	postgresDatabase *pgxpool.Pool
}

//factory

func NewWorkspaceUserRepository(posgresDatabase *pgxpool.Pool) WorkspaceUserRepository {
	return WorkspaceUserRepository{postgresDatabase: posgresDatabase}
}

//methods

func (w *WorkspaceUserRepository) Create(userId string, nickname string) (*WorkspaceUser, error) {
	tx, err := w.postgresDatabase.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	generatedId := uuid.NewString()
	_, err = tx.Exec(context.Background(), `INSERT INTO workspace_user (id, user_id, nickname) VALUES ($1, $2, $3)`, generatedId, userId, nickname)
	if err != nil {
		return nil, err
	}

	workspaceUser := WorkspaceUser{Id: generatedId, UserId: userId, Nickname: nickname}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}
	return &workspaceUser, nil
}

func (w *WorkspaceUserRepository) Read(id string) (*WorkspaceUser, error) {
	row, err := w.postgresDatabase.Query(context.Background(), `SELECT * FROM workspace_user WHERE user_id=$1`, id)
	if err != nil {
		return nil, err
	}

	var user WorkspaceUser
	err = pgxscan.ScanOne(&user, row)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
