import { Module } from '@nestjs/common'
import { AppController } from '@/app.controller'
import { AuthModule } from '@/modules/auth/auth.module'
import { UserModule } from '@/modules/user/user.module'
import { AnimeModule } from '@/modules/anime/anime.module'
import { ConfigModule } from '@nestjs/config'
import { EnvConfigOptions } from '@/configs'

@Module({
	imports: [AuthModule, UserModule, AnimeModule, ConfigModule.forRoot(EnvConfigOptions)],
	controllers: [AppController]
})
export class AppModule {}
