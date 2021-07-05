import { INestApplication, Injectable } from '@nestjs/common';

@Injectable()
export class WorkspaceUserService {

    constructor(private workspaceUserRepository: WorkspaceUserService) {
    }

    createWorkspaceUser() {

    }
}
