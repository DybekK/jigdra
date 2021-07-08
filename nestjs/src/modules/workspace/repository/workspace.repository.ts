import {PrismaRepository} from "../../../database/repository";
import {Prisma} from "@prisma/client";

export class WorkspaceRepository extends PrismaRepository {
    get prisma(): Prisma.WorkspaceDelegate<any> {
        return this.em.workspace;
    }
}