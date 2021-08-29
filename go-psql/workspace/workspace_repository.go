package workspace

import "github.com/jackc/pgx/v4/pgxpool"

type WorkspaceRepository struct {
	postgresDatabase *pgxpool.Pool
}

//factory

func NewWorkspaceRepository(postgresDatabase *pgxpool.Pool) WorkspaceRepository {
	return WorkspaceRepository{postgresDatabase: postgresDatabase}
}
