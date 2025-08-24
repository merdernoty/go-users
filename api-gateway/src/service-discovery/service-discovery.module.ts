import { Module } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'

@Module({
	providers: [ServiceDiscovery],
	exports: [ServiceDiscovery]
})
export class ServiceDiscoveryModule {}
