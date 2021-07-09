import { IsNotEmpty, IsString } from "class-validator";

export class CreateWorkspaceUserDto {
    @IsString()
    @IsNotEmpty()
    userId: string;

    @IsString()
    @IsNotEmpty()
    nickname: string;
}