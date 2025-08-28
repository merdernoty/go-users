import { NestFactory } from '@nestjs/core'
import { AppModule } from './app.module'
import { clusterModeBootstrap } from './cluster'
import { ServiceDiscoveryModule } from './service-discovery/module'
import { ConfigService } from '@nestjs/config'

async function bootstrap() {
	const app = await NestFactory.create(AppModule)

	app.setGlobalPrefix('api')

	const configService = app.get<ConfigService>('ConfigService')
	const PORT = configService.get('PORT')

	await app.listen(PORT)
}

async function discoveryServiceBootstrap() {
	const app = await NestFactory.create(ServiceDiscoveryModule)
	const configService = app.get<ConfigService>('ConfigService')
	const PORT = configService.get('PORT')

	await app.listen(PORT)
}

const main = async () => {
	const isClusterMode = process.env.CLUSTER_MODE === 'true'

	if (isClusterMode) {
		clusterModeBootstrap(bootstrap, discoveryServiceBootstrap)
		return
	} else {
		bootstrap()
		discoveryServiceBootstrap()
	}
}

main()
