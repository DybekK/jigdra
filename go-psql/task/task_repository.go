package task

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TaskRepository struct {
	postgresDatabase *pgxpool.Pool
}

//factory
func NewTaskRepository(postgresDatabase *pgxpool.Pool) TaskRepository {
	return TaskRepository{postgresDatabase: postgresDatabase}
}

//methods
func (r *TaskRepository) Create(task Task) (*Task, error) {
	tx, err := r.postgresDatabase.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())
	query := `INSERT INTO task (id, workspace_user_id, workspace_id, title) VALUES ($1,$2,$3,$4)`
	_, err = tx.Exec(context.Background(), query, task.Id, task.WorkspaceUserId, task.WorkspaceId, task.Title)
	if err != nil {
		fmt.Println()
		return nil, err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) ReadByTaskId(task_id string) (*Task, error) {
	query := `SELECT * FROM task WHERE id=$1`
	row, err := r.postgresDatabase.Query(context.Background(), query, task_id)
	if err != nil {
		return nil, err
	}
	var task Task
	err = pgxscan.ScanOne(&task, row)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) ReadByUserId(user_id string) ([]Task, error) {
	query := `SELECT * FROM task WHERE workspace_user_id=$1`
	rows, err := r.postgresDatabase.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = pgxscan.ScanAll(&tasks, rows)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
