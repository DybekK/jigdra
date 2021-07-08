import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { EntityManager } from './database/entityManager.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: '.env',
      isGlobal: true,
    }),
  ],
  controllers: [],
  providers: [EntityManager],
})
export class AppModule {}
