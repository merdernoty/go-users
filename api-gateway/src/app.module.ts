import { Module } from '@nestjs/common'
import { AppController } from '@/app.controller'
import { AuthModule } from '@/modules/auth/auth.module'
import { UserModule } from '@/modules/user/user.module'
import { AnimeModule } from './modules/anime/anime.module'

@Module({
	imports: [AuthModule, UserModule, AnimeModule],
	controllers: [AppController]
})
export class AppModule {}
