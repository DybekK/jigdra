import {Injectable} from "@nestjs/common";
import {InjectRepository} from "@nestjs/typeorm";
import {Task} from "../../../../database/entity/entity";
import {Repository} from "typeorm";
import {TaskRepository} from "../../repository/task.repository";
import {FindTaskByCriteriaDto} from "../../dto/findTaskByCriteriaDto";
import {CreateTaskDto} from "../../dto/createTaskDto";

@Injectable()
export class TaskService {
    constructor(
        @InjectRepository(TaskRepository)
        private taskRepository: TaskRepository
    ) {}

    async getTasks({workspaceId, workspaceUserId}: FindTaskByCriteriaDto): Promise<Task[]> {
        return this.taskRepository.findByCriteria({workspaceId, workspaceUserId});
    }

    async createTask(createTaskDto: CreateTaskDto) {
        const task = new Task();
        return await this.taskRepository.save(task);
    }
}