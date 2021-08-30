package workspace

type WorkspaceFacade struct {
	workspaceService     WorkspaceService
	workspaceUserService WorkspaceUserService
}

//factory

func NewWorkspaceFacade(workspaceService WorkspaceService, workspaceUserService WorkspaceUserService) WorkspaceFacade {
	return WorkspaceFacade{workspaceService: workspaceService, workspaceUserService: workspaceUserService}
}

func (w *WorkspaceFacade) CreateUserAndWorkspace(userId string, nickname string) (*WorkspaceUser, *Workspace) {
	workspaceUser, _ := w.workspaceUserService.CreateUser(userId, nickname)
	workspace, _ := w.workspaceService.CreateWorkspace(workspaceUser.Id)
	return workspaceUser, workspace
}
