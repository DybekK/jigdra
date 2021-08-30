package workspace

type WorkspaceFacade struct {
	workspaceService     WorkspaceService
	workspaceUserService WorkspaceUserService
}

//factory

func NewWorkspaceFacade(workspaceService WorkspaceService, workspaceUserService WorkspaceUserService) WorkspaceFacade {
	return WorkspaceFacade{workspaceService: workspaceService, workspaceUserService: workspaceUserService}
}

//methods

func (w *WorkspaceFacade) CreateUserAndWorkspace(userId string, nickname string) (*WorkspaceUser, *Workspace, error) {
	workspace, err := w.workspaceService.CreateWorkspace()
	if err != nil {
		return nil, nil, err
	}
	workspaceUser, err := w.workspaceUserService.CreateUser(userId, workspace.Id, nickname)
	if err != nil {
		return nil, nil, err
	}
	return workspaceUser, workspace, err
}
