import {Body, Controller, Post} from "@nestjs/common";
import {WorkspaceFacadeService} from "../service/workspace-facade/workspace-facade.service";
import {CreateWorkspaceDto} from "../dto/createWorkspaceDto";

@Controller("api")
export class WorkspaceController {
    constructor(private workspaceServiceFacade: WorkspaceFacadeService) {}

    @Post("workspace")
    async createWorkspaceForUser(@Body() createWorkspaceDto: CreateWorkspaceDto) {
        return this.workspaceServiceFacade.createWorkspaceForUser(createWorkspaceDto);
    }
}