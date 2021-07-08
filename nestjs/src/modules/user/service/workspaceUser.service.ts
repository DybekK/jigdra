import { Injectable } from '@nestjs/common';
import { WorkspaceUser } from '@prisma/client';
import {WorkspaceUserRepository} from "../repository/workspaceUser.repository";

@Injectable()
export class WorkspaceUserService {
    constructor(private workspaceUserRepository: WorkspaceUserRepository) {}
}
