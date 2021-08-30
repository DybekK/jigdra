package workspace

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WorkspaceRepository struct {
	postgresDatabase *pgxpool.Pool
}

//factory

func NewWorkspaceRepository(postgresDatabase *pgxpool.Pool) WorkspaceRepository {
	return WorkspaceRepository{postgresDatabase: postgresDatabase}
}

//methods

func (w *WorkspaceRepository) Create() (*Workspace, error) {
	tx, err := w.postgresDatabase.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	generatedId := uuid.NewString()
	_, err = tx.Exec(context.Background(), `INSERT INTO workspace (id) VALUES ($1)`, generatedId)
	if err != nil {
		return nil, err
	}

	workspace := Workspace{Id: generatedId}
	err = tx.Commit(context.Background())
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (w *WorkspaceRepository) Read(id string) (*Workspace, error) {
	row, err := w.postgresDatabase.Query(context.Background(), `SELECT * FROM workspace WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}

	var workspace Workspace
	err = pgxscan.ScanOne(&workspace, row)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}
