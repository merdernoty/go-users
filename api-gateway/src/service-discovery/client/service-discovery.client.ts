import { ConfigService } from '@nestjs/config'
//TODO:  Можно заменить на http module из nestjs
import axios from 'axios'
import { ClientNameKey, clientMap } from '../clients.constant'
import { ClientGrpcProxy, ClientProxyFactory, GrpcOptions } from '@nestjs/microservices'

export class ServiceDiscoveryClient {
	constructor(private readonly configService: ConfigService) {}

	public async getClient(name: ClientNameKey): Promise<{ isSuccess: boolean; client?: ClientGrpcProxy }> {
		try {
			const port = this.configService.get<string>('SERVICE_DISCOVERY_PORT')
			const url = `http://localhost:${port}`
			const result = await axios.get(`${url}/service-discovery/${name}/host`)


            const client = this.createClient(name, result.data.host)
			return {
				isSuccess: true,
				client
			}
		} catch (err) {
			return {
				isSuccess: false
			}
		}
	}

	private createClient(name: ClientNameKey, url: string) {
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
}
