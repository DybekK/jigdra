package workspace

type WorkspaceFacade struct {
	workspaceService     WorkspaceService
	workspaceUserService WorkspaceUserService
}

//factory

func NewWorkspaceFacade(workspaceService WorkspaceService, workspaceUserService WorkspaceUserService) WorkspaceFacade {
	return WorkspaceFacade{workspaceService: workspaceService, workspaceUserService: workspaceUserService}
}
