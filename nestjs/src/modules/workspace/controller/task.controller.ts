import {Body, Controller, Get, Post, Req} from "@nestjs/common";
import {TaskService} from "../service/task/task.service";
import {FindTaskByCriteriaDto} from "../dto/findTaskByCriteriaDto";
import {Request} from "express";
import {CreateTaskDto} from "../dto/createTaskDto";

@Controller("api")
export class TaskController {
    constructor(private taskService: TaskService) {}

    @Get("task")
    async getTasks(@Req() req: Request) {
        return this.taskService.getTasks({...req.params} as FindTaskByCriteriaDto);
    }

    @Post("task")
    async createTask(@Body() createTaskDto: CreateTaskDto) {
        return this.taskService.createTask(createTaskDto);
    }
}