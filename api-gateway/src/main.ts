import { NestFactory } from '@nestjs/core'
import { AppModule } from './app.module'
import { clusterModeBootstrap } from './cluster'
import { ServiceDiscoveryModule } from './service-discovery'

async function bootstrap() {
	const app = await NestFactory.create(AppModule)
	await app.listen(process.env.PORT ?? 3000)
}

async function discoveryServiceBootstrap() {
	const app = await NestFactory.create(ServiceDiscoveryModule)
	await app.listen(process.env.PORT ?? 3001)
}

const main = async () => {
	const isClusterMode = process.env.CLUSTER_MODE === 'true'

	if(isClusterMode) {
		clusterModeBootstrap(bootstrap, discoveryServiceBootstrap)
		return;
	}else {
		bootstrap();
		discoveryServiceBootstrap();
	}
}

main()
