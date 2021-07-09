import {InjectRepository} from "@nestjs/typeorm";
import {WorkspaceUser} from "../../../../database/entity/entity";
import {Repository} from "typeorm";
import {CreateWorkspaceUserDto} from "../../dto/createWorkspaceUserDto";
import {Injectable} from "@nestjs/common";

@Injectable()
export class WorkspaceUserService {
    constructor(
        @InjectRepository(WorkspaceUser)
        private workspaceUserRepository: Repository<WorkspaceUser>
    ) {}

    async createWorkspaceUser(createWorkspaceUserDto: CreateWorkspaceUserDto): Promise<WorkspaceUser> {
        const workspaceUser = new WorkspaceUser();
        workspaceUser.userId = createWorkspaceUserDto.userId;
        workspaceUser.nickname = createWorkspaceUserDto.nickname;
        return this.workspaceUserRepository.save(workspaceUser);
    }
}