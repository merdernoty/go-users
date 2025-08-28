import { Controller, Get, Param, ParseIntPipe } from '@nestjs/common'
import { GetUserResponse, UserService } from '@grpc-types/user/user'
import { ServiceDiscoveryClient } from '@/service-discovery/client'

@Controller('user')
export class UserController {
	constructor(private serviceDiscoveryClient: ServiceDiscoveryClient) {}

	private async userService(): Promise<UserService> {
		const result = await this.serviceDiscoveryClient.getClient('USER_CLIENT')
		if (!result.isSuccess || !result.client) {
			throw new Error('Anime service is unavailable')
		}

		return result.client.getService<UserService>('UserService')
	}

	@Get(':id')
	public async getOne(@Param('id', ParseIntPipe) id: Long): Promise<GetUserResponse> {
		const userService = await this.userService()
		return userService.GetUser({ id })
	}

	@Get(':email')
	public async getUserByEmail() {
		const userService = await this.userService()
		return userService.GetUserByEmail({ email: '' })
	}

	@Get()
	public async getAll() {
		const userService = await this.userService()
		return userService.ListUsers({})
	}
}
