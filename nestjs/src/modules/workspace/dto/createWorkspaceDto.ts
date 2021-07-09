import { IsNotEmpty, IsString } from 'class-validator';
import {CreateWorkspaceUserDto} from "./createWorkspaceUserDto";


export class CreateWorkspaceDto extends CreateWorkspaceUserDto {
    @IsString()
    @IsNotEmpty()
    name: string;
}