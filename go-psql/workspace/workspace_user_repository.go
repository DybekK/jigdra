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

func (w *WorkspaceUserRepository) Create(userId string, workspaceId string, nickname string) (*WorkspaceUser, error) {
	tx, err := w.postgresDatabase.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	generatedId := uuid.NewString()
	query := "INSERT INTO workspace_user (id, user_id, workspace_id, nickname) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(context.Background(), query, generatedId, userId, workspaceId, nickname)
	if err != nil {
		return nil, err
	}

	workspaceUser := WorkspaceUser{Id: generatedId, UserId: userId, WorkspaceId: workspaceId, Nickname: nickname}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}
	return &workspaceUser, nil
}

func (w *WorkspaceUserRepository) ReadByMongoId(mongo_id string) (*WorkspaceUser, error) {
	row, err := w.postgresDatabase.Query(context.Background(), `SELECT * FROM workspace_user WHERE user_id=$1`, mongo_id)
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

func (w *WorkspaceUserRepository) Read(id string) (*WorkspaceUser, error) {
	row, err := w.postgresDatabase.Query(context.Background(), `SELECT * FROM workspace_user WHERE id=$1`, id)
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
