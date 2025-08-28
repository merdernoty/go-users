import { Injectable } from '@nestjs/common'
import { ServiceDiscovery } from './service-discovery'
import { CronExpression, SchedulerRegistry } from '@nestjs/schedule'
import { ConfigService } from '@nestjs/config'
import { CronJob } from 'cron'

@Injectable()
export class ServiceDiscoveryController {
	private readonly cronJobName = 'service-health-check'

	constructor(
		private readonly configService: ConfigService,
		private readonly serviceDiscovery: ServiceDiscovery,
		private readonly schedulerRegistry: SchedulerRegistry
	) {
		const cronTimeHealthCheck = this.configService.get<string>(
			'CRON_TIME_HEALTH_CHECK',
			CronExpression.EVERY_5_MINUTES
		)

        this.startJob(cronTimeHealthCheck)
	}

	private startJob(time: string) {
		const job = new CronJob(time, this.handleCron)
		this.schedulerRegistry.addCronJob(this.cronJobName, job)
		job.start()
		// this.logger.warn(`job ${name} added for each minute at ${seconds} seconds!`)
	}

	private async handleCron() {
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
