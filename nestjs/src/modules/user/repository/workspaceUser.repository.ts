import {Injectable} from "@nestjs/common";
import {Prisma, WorkspaceUser } from "@prisma/client";
import {PrismaRepository} from "../../../database/repository";

@Injectable()
export class WorkspaceUserRepository extends PrismaRepository {
    get prisma(): Prisma.WorkspaceUserDelegate<any> {
        return this.em.workspaceUser;
    }
}