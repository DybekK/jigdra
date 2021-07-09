import {Workspace, WorkspaceUser } from "@prisma/client";

export interface WorkspaceForUser {
    workspaceUser: WorkspaceUser,
    workspace: Workspace
}