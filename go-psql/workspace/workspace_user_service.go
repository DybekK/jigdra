package workspace

import "github.com/google/uuid"

type WorkspaceUserService struct {
	workspaceUserRepository WorkspaceUserRepository
}

//factory

func NewWorkspaceUserService(workspaceUserRepository WorkspaceUserRepository) WorkspaceUserService {
	return WorkspaceUserService{workspaceUserRepository: workspaceUserRepository}
}

//methods

func (wus *WorkspaceUserService) GetUser(id string) (*WorkspaceUser, error) {
	return wus.workspaceUserRepository.Read(id)
}

func (wus *WorkspaceUserService) CreateUser(mongo_id, nickname string) error {
	var user WorkspaceUser
	user.Id = uuid.New()
	user.UserId = mongo_id
	user.Nickname = nickname
	return wus.workspaceUserRepository.Create(user)
}
