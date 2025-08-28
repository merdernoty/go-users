import { Controller, Get, Post } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'
import { Cron, CronExpression } from '@nestjs/schedule'

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

	@Cron(CronExpression.EVERY_5_MINUTES)
	async handleCron() {
		try {
		} catch (error) {}

		const info = this.serviceDiscovery.clientsInfo
		const services = info.keys()

		const newInfo: { name: string; host: string }[] = []

		for (const name of services) {
			const result = await this.serviceDiscovery.pingService(name)

			const hosts = Object.keys(result).filter(host => result[host])

			hosts.forEach(host => {
				newInfo.push({ name, host })
			})

			console.log(`Health check for service ${name}:`, JSON.stringify(result, null, 2))
		}


        this.serviceDiscovery.healthCheck(newInfo)
	}
}
