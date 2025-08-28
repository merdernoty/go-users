import { Module } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'
import { ServiceDiscoveryController } from './service-discovery.controller'
import { ScheduleModule } from '@nestjs/schedule'

@Module({
	imports: [ScheduleModule.forRoot()],
	controllers: [ServiceDiscoveryController],
	providers: [ServiceDiscovery],
	exports: [ServiceDiscovery]
})
export class ServiceDiscoveryModule {}
