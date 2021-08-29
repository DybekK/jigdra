package workspace

type WorkspaceService struct {
	workspaceRepository WorkspaceRepository
}

//factory

func NewWorkspaceService(workspaceRepository WorkspaceRepository) WorkspaceService {
	return WorkspaceService{workspaceRepository: workspaceRepository}
}
