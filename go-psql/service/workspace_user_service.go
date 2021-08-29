package service

import (
	"go-psql/database/repository"
	"go-psql/dto"

	"github.com/google/uuid"
)

type WorkspaceUserService struct {
	workspaceUserRepository repository.WorkspaceUserRepository
}

//factory

func NewWorkspaceUserService(workspaceUserRepository repository.WorkspaceUserRepository) WorkspaceUserService {
	return WorkspaceUserService{workspaceUserRepository: workspaceUserRepository}
}

//methods

func (wus *WorkspaceUserService) GetUser(id string) *dto.WorkspaceUser {
	return wus.workspaceUserRepository.Read(id)
}

func (wus *WorkspaceUserService) CreateUser(mongo_id, nickname string) error {
	var user dto.WorkspaceUser
	user.Id = uuid.New()
	user.UserId = mongo_id
	user.Nickname = nickname
	return wus.workspaceUserRepository.Create(user)
}
