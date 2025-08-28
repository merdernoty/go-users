import { Injectable } from '@nestjs/common'
import axios from 'axios'

@Injectable()
export class ServiceDiscovery {
	public clientsInfo: Map<string, string[]> = new Map()

	constructor() {}

	// TODO: Возможно вынести балансировку в отдельный сервис
	public getHost(name: string): string {
		const clientInfo = this.clientsInfo.get(name)

		if (!clientInfo || clientInfo.length === 0) {
			// TODO: Продумать обработку этой ошибки
			throw new Error(`Service ${name} not available`)
		}

		//TODO: Балансировка нагрузки, продумать другую стратегию
		const url = clientInfo[Math.floor(Math.random() * clientInfo.length)]
		return url
	}

	public addClient(name: string, host: string) {
		const clientInfo = this.clientsInfo.get(name) || []
		clientInfo.push(host)
		this.clientsInfo.set(name, clientInfo)
	}

	private async ping(host: string): Promise<boolean> {
		try {
			const response = await axios.get(`${host}/ping`)
			return response.status === 200
		} catch (error) {
			return false
		}
	}

	public async pingService(name: string) {
		const hosts = this.clientsInfo.get(name) || []

		const result: Record<string, boolean> = {}

		for (const host of hosts) {
			const isAlive = await this.ping(host)
			result[host] = isAlive
		}
		return result;
	}

	// TODO: Продумать модель взаимодействия с ServiceDiscovery (CRON, SSE, pull/push, sub/pub)
	public healthCheck(info: { name: string; host: string }[]) {
		this.clientsInfo.clear()
		const serviceName = info.map(item => item.name)

		serviceName.forEach(name => {
			this.clientsInfo.set(name, [])
		})

		info.forEach(item => {
			const list = this.clientsInfo.get(item.name) || []
			list.push(item.host)
			this.clientsInfo.set(item.name, list)
		})
	}
}

// CRON /ping -> сервисы
// services: auth, anime, users
