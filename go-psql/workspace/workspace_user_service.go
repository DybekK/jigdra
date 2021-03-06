package workspace

type WorkspaceUserService struct {
	workspaceUserRepository WorkspaceUserRepository
}

//factory

func NewWorkspaceUserService(workspaceUserRepository WorkspaceUserRepository) WorkspaceUserService {
	return WorkspaceUserService{workspaceUserRepository: workspaceUserRepository}
}

//methods

func (w *WorkspaceUserService) GetUserByMongoId(mongo_id string) (*WorkspaceUser, error) {
	return w.workspaceUserRepository.ReadByMongoId(mongo_id)
}

func (w *WorkspaceUserService) GetUser(id string) (*WorkspaceUser, error) {
	return w.workspaceUserRepository.Read(id)
}

func (w *WorkspaceUserService) CreateUser(userId string, workspaceId string, nickname string) (*WorkspaceUser, error) {
	return w.workspaceUserRepository.Create(userId, workspaceId, nickname)
}
