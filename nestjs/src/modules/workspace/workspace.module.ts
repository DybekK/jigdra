import {Module} from "@nestjs/common";
import {TypeOrmModule} from "@nestjs/typeorm";
import {Workspace, WorkspaceUser} from "../../database/entity/entity";
import {WorkspaceService} from "./service/workspace/workspace.service";
import {WorkspaceController} from "./controller/workspace.controller";
import {WorkspaceUserService} from "./service/workspace-user/workspace-user.service";
import {WorkspaceFacadeService} from "./service/workspace-facade/workspace-facade.service";
import {EntityManager} from "typeorm";

@Module({
    controllers: [WorkspaceController],
    imports: [TypeOrmModule.forFeature([Workspace, WorkspaceUser])],
    providers: [WorkspaceUserService, WorkspaceService, WorkspaceFacadeService]
})
export class WorkspaceModule {

}