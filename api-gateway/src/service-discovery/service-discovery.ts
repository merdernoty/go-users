import { Injectable } from '@nestjs/common'
import { ClientGrpcProxy, ClientProxyFactory, GrpcOptions } from '@nestjs/microservices'
import { ClientNameKey, clientMap } from './clients.constant'

@Injectable()
export class ServiceDiscovery {
	private clientsInfo: Map<string, string[]> = new Map()

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

	public getClient(name: ClientNameKey): ClientGrpcProxy {
		//TODO: Возможно оставить ошибку тут или продумать обработку
		const url = this.getHost(name)
		const config = clientMap[name]

		const option: GrpcOptions = {
			...config,
			options: {
				...config.options,
				url
			}
		}

		const client = ClientProxyFactory.create(option) as ClientGrpcProxy
		return client
	}

	// * TODO: Продумать модель взаимодействия с ServiceDiscovery (CRON, SSE, pull/push, sub/pub)
	private healthCheck(info: { name: string; host: string }[]) {
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
