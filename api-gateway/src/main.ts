import { NestFactory } from '@nestjs/core'
import { AppModule } from './app.module'
import { clusterModeBootstrap } from './cluster'

async function bootstrap() {
	const app = await NestFactory.create(AppModule)
	await app.listen(process.env.PORT ?? 3000)
}

const main = async () => {
	const isClusterMode = process.env.CLUSTER_MODE === 'true'
	if (isClusterMode) {
		clusterModeBootstrap(bootstrap)
	} else {
		bootstrap()
	}
}

main()
