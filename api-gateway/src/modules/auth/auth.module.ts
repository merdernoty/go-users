import { Module } from '@nestjs/common'
import { AuthController } from './auth.controller'
import { ServiceDiscoveryModule } from '@/service-discovery'

@Module({
	imports: [ServiceDiscoveryModule],
	controllers: [AuthController]
})
export class AuthModule {}
