import { AuthService } from '@/generated/auth/auth'
import { ServiceDiscovery } from '@/service-discovery'
import { Body, Controller, Post } from '@nestjs/common'
import { LoginAuthDto, RegisterAuthDto } from './dto'

@Controller('auth')
export class AuthController {
	constructor(private serviceDiscovery: ServiceDiscovery) {}

	private get authService(): AuthService {
		const client = this.serviceDiscovery.getClient('AUTH_CLIENT')
		return client.getService<AuthService>('AuthService')
	}

	@Post('login')
	public async login(@Body() dto: LoginAuthDto) {
		const response = await this.authService.Login(dto)
		return response
	}

	@Post('register')
	public register(@Body() dto: RegisterAuthDto) {
		const response = this.authService.Register(dto)
		return response
	}
}
