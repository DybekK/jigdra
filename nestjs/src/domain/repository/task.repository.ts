import { Repository } from './repository';
import { Injectable } from '@nestjs/common';

@Injectable()
export class TaskRepository extends Repository {}
