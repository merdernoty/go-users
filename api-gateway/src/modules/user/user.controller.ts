import { ServiceDiscovery } from '@/service-discovery'
import { Controller, Get, Param, ParseIntPipe } from '@nestjs/common'
import { GetUserResponse, UserService } from '@grpc-types/user/user'

@Controller('user')
export class UserController {
	constructor(private serviceDiscovery: ServiceDiscovery) {}

	private get userService(): UserService {
		const client = this.serviceDiscovery.getClient('USER_CLIENT')
		return client.getService<UserService>('UserService')
	}

	@Get(':id')
	public async getOne(@Param('id', ParseIntPipe) id: Long): Promise<GetUserResponse> {
		return this.userService.GetUser({ id })
	}

	@Get(':email')
	public getUserByEmail() {
		return this.userService.GetUserByEmail({ email: '' })
	}

	@Get()
	public getAll() {
		return this.userService.ListUsers({})
	}
}
