import { Module } from '@nestjs/common';
import {WorkspaceService} from "../workspace/service/workspace/workspace.service";

@Module({
    exports: [
      WorkspaceService
    ],
    providers: [
        WorkspaceService
    ]
})
export class UserModule {}
