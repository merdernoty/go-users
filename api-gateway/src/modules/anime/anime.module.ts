import { ServiceDiscoveryModule } from '@/service-discovery'
import { Module } from '@nestjs/common'
import { AnimeController } from './anime.controller'

@Module({
	imports: [ServiceDiscoveryModule],
	controllers: [AnimeController]
})
export class AnimeModule {}
