import {Injectable} from "@nestjs/common";
import {WorkspaceService} from "../workspace/workspace.service";
import {WorkspaceUserService} from "../workspace-user/workspace-user.service";
import {CreateWorkspaceUserDto} from "../../dto/createWorkspaceUserDto";
import {CreateWorkspaceDto} from "../../dto/createWorkspaceDto";
import {EntityManager} from "typeorm";
import {WorkspaceUser} from "../../../../database/entity/entity";

@Injectable()
export class WorkspaceFacadeService {
    constructor(
        private workspaceService: WorkspaceService,
        private workspaceUserService: WorkspaceUserService,
        private em: EntityManager
    ) {}

    async createWorkspaceForUser(createWorkspaceDto: CreateWorkspaceDto): Promise<WorkspaceUser> {
        const workspace = await this.workspaceService.createWorkspace(createWorkspaceDto);
        const workspaceUser = await this.workspaceUserService.createWorkspaceUser(createWorkspaceDto);
        workspaceUser.workspace = Promise.resolve(workspace);
        await this.em.save(workspaceUser);
        await this.em.save(workspace);
        return workspaceUser;
    }
}