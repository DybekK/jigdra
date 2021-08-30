package workspace

type WorkspaceService struct {
	workspaceRepository WorkspaceRepository
}

//factory

func NewWorkspaceService(workspaceRepository WorkspaceRepository) WorkspaceService {
	return WorkspaceService{workspaceRepository: workspaceRepository}
}

func (w *WorkspaceService) GetWorkspace(id string) (*Workspace, error) {
	return w.workspaceRepository.Read(id)
}

func (w *WorkspaceService) CreateWorkspace(workspaceUserId string) (*Workspace, error) {
	return w.workspaceRepository.Create(workspaceUserId)
}
