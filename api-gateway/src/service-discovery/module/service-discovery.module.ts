import { Module } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'
import { ServiceDiscoveryController } from './service-discovery.controller'
import { ScheduleModule } from '@nestjs/schedule'
import { ConfigModule } from '@nestjs/config'
import { EnvServiceDiscoveryConfigOptions } from '@/configs'

@Module({
	imports: [ScheduleModule.forRoot(), ConfigModule.forRoot(EnvServiceDiscoveryConfigOptions)],
	controllers: [ServiceDiscoveryController],
	providers: [ServiceDiscovery],
	exports: [ServiceDiscovery]
})
export class ServiceDiscoveryModule {}
