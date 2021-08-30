package workspace

type WorkspaceService struct {
	workspaceRepository WorkspaceRepository
}

//factory

func NewWorkspaceService(workspaceRepository WorkspaceRepository) WorkspaceService {
	return WorkspaceService{workspaceRepository: workspaceRepository}
}

//methods

func (w *WorkspaceService) GetWorkspace(id string) (*Workspace, error) {
	return w.workspaceRepository.Read(id)
}

func (w *WorkspaceService) CreateWorkspace() (*Workspace, error) {
	return w.workspaceRepository.Create()
}
