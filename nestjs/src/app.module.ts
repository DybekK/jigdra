import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import {TypeOrmModule} from "@nestjs/typeorm";
import {Workspace, WorkspaceUser} from "./database/entity/entity";
import {WorkspaceModule} from "./modules/workspace/workspace.module";
import {WorkspaceController} from "./modules/workspace/controller/workspace.controller";

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env',
      isGlobal: true,
    }),
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: 'localhost',
      port: 5432,
      database: 'jidgra',
      username: 'admin',
      password: 'passwd',
      entities: [Workspace, WorkspaceUser],
      synchronize: true
    }),
    WorkspaceModule
  ],
})
export class AppModule {}
