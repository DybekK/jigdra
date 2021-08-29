package workspace

type WorkspaceUserService struct {
	workspaceUserRepository WorkspaceUserRepository
}

//factory

func NewWorkspaceUserService(workspaceUserRepository WorkspaceUserRepository) WorkspaceUserService {
	return WorkspaceUserService{workspaceUserRepository: workspaceUserRepository}
}

//methods

func (wus *WorkspaceUserService) GetUser(id string) *WorkspaceUser {
	return wus.workspaceUserRepository.Read(id)
}
