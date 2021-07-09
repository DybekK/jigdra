import {Injectable} from "@nestjs/common";
import {InjectRepository} from "@nestjs/typeorm";
import {Workspace} from "../../../../database/entity/entity";
import {Repository} from "typeorm";
import {CreateWorkspaceDto} from "../../dto/createWorkspaceDto";

@Injectable()
export class WorkspaceService {
    constructor(
        @InjectRepository(Workspace)
        private workspaceRepository: Repository<Workspace>,
    ) {}

    async createWorkspace(createWorkspaceDto: CreateWorkspaceDto) {
        const workspace = new Workspace();
        workspace.name = createWorkspaceDto.name;
        return this.workspaceRepository.save(workspace);
    }
}