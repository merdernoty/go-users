import { AuthService } from '@/generated/auth/auth'
import { Body, Controller, Post } from '@nestjs/common'
import { LoginAuthDto, RegisterAuthDto } from './dto'
import { ServiceDiscoveryClient } from '@/service-discovery/client'

@Controller('auth')
export class AuthController {
	constructor(private serviceDiscoveryClient: ServiceDiscoveryClient) {}

	private async authService(): Promise<AuthService> {
		const result = await this.serviceDiscoveryClient.getClient('AUTH_CLIENT')
		if (!result.isSuccess || !result.client) {
			throw new Error('Anime service is unavailable')
		}

		return result.client.getService<AuthService>('AuthService')
	}

	@Post('login')
	public async login(@Body() dto: LoginAuthDto) {
		const authService = await this.authService()
		const response = await authService.Login(dto)
		return response
	}

	@Post('register')
	public async register(@Body() dto: RegisterAuthDto) {
		const authService = await this.authService()
		const response = authService.Register(dto)
		return response
	}
}
