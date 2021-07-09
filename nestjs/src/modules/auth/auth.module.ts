import { Module } from '@nestjs/common';
import { AuthGuard } from './guard/auth.guard';

@Module({
  exports: [AuthGuard],
  providers: [AuthGuard],
})
export class AuthModule {}
