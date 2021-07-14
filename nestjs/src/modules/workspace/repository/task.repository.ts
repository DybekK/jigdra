import {EntityRepository, Repository} from "typeorm";
import {Task} from "../../../database/entity/entity";
import {FindTaskByCriteriaDto} from "../dto/findTaskByCriteriaDto";

@EntityRepository(Task)
export class TaskRepository extends Repository<Task> {
    async findByCriteria(criteria: FindTaskByCriteriaDto): Promise<Task[]> {
        const qb = this.createQueryBuilder("task")
            .select("task");

        if(criteria.workspaceId) {
            qb.leftJoin("task.workspace", "workspace")
                .where("workspace.id = :id", { id: criteria.workspaceId })
        }

        if(criteria.workspaceUserId) {
            qb.leftJoin("task.workspaceUsers", "workspaceUser")
                .where("workspaceUser.id = :id", { id: criteria.workspaceUserId })
        }

        return qb.getMany()
    }
}