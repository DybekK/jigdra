import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EntityManager } from './domain/entityManager.service';
import {TestController} from "./controller/test.controller";

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env',
      isGlobal: true,
    }),
  ],
  controllers: [TestController],
  providers: [EntityManager],
})
export class AppModule {}
