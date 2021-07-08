import { Injectable } from '@nestjs/common';
import {WorkspaceRepository} from "../../repository/workspace.repository";
import {WorkspaceUserRepository} from "../../../user/repository/workspaceUser.repository";

@Injectable()
export class WorkspaceService {
    constructor(
        private workspaceUserRepository: WorkspaceUserRepository,
        private workspaceRepository: WorkspaceRepository
    ) {}

    async createWorkspaceForUser(workspaceUserId: string) {
        
    }
}
