import { Controller, Get, Post } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'

@Controller('service-discovery')
export class ServiceDiscoveryController {
	constructor(private readonly serviceDiscovery: ServiceDiscovery) {}

	@Get(':name/host')
	async getClient(name: string) {
		const host = this.serviceDiscovery.getHost(name)
		return { host }
	}

	@Post('add')
	async addClient(name: string, host: string) {
		this.serviceDiscovery.addClient(name, host)
		return { message: 'Client added successfully' }
	}
}
