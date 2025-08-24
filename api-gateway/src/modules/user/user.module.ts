import { ServiceDiscoveryModule } from '@/service-discovery'
import { Module } from '@nestjs/common'
import { UserController } from './user.controller'

@Module({
	imports: [ServiceDiscoveryModule],
	controllers: [UserController]
})
export class UserModule {}
