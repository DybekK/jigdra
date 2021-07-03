import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import {ConfigModule} from "@nestjs/config";
import configuration from "./config/configuration";

const {database} = configuration();

@Module({
  imports: [
      ConfigModule.forRoot({
        load: [configuration],
        envFilePath: '.env',
        isGlobal: true
      }),
      SequelizeModule.forRoot({
        dialect: 'postgres',
        host: database.host,
        port: database.port,
        username: database.username,
        password: database.password,
        database: database.name,
        models: [],
      })
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
