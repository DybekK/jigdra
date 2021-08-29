package service

import (
	"go-psql/database/repository"
	"go-psql/dto"
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
