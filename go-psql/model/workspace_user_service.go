package model

import (
	"fmt"
	"go-psql/model/dto"
	"go-psql/model/repository"

	"github.com/google/uuid"
)

type wus struct {
	repo repository.WorkspaceUserRepo
}

type WorkspaceUserService interface {
	SaveWorkspaceUser(string, string) (*uuid.UUID, error)
	FindByUserId(string) *dto.WorkspaceUser
}

func NewWorkspaceUserService(repo repository.WorkspaceUserRepo) WorkspaceUserService {
	return &wus{repo: repo}
}

func (s *wus) SaveWorkspaceUser(mongo_id, nickname string) (*uuid.UUID, error) {
	var user dto.WorkspaceUser
	user.Id = uuid.New()
	user.UserId = mongo_id
	user.Nickname = nickname
	fmt.Println(user.Nickname)
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return &user.Id, nil
}

func (s *wus) FindByUserId(mongo_id string) *dto.WorkspaceUser {
	user := s.repo.Read(mongo_id)
	return user
}
