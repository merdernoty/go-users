import cluster from 'cluster'
import { cpus } from 'os'


// TODO: Проблема в Cluster mode, 
// TODO: в том что не будет общий Discovery Service -> либо в отдельный воркер или в хранить записи 
// TODO: в Redis/memchached если возможно
export const clusterModeBootstrap = (bootstrap: () => Promise<void> | void, discoveryServiceBootstraps: () => Promise<void> | void) => {
	if (cluster.isPrimary) {
		const numCPUs = cpus().length - 2

		
		for (let i = 0; i < numCPUs; i++) {
			cluster.fork({...process.env, isWorker: true, isServiceDiscovery: false})
		}

		cluster.fork({...process.env, isWorker: false, isServiceDiscovery: true})
	

		cluster.on('exit', (worker, code, signal) => {
			console.log(`worker ${worker.process.pid} code: ${code} died`);
			cluster.fork({...process.env, isWorker: true, isServiceDiscovery: false})
			console.log(`worker ${worker.process.pid} restarted`);
		})
	} else {
		const {isWorker, isServiceDiscovery} = process.env

		if(isWorker) {
			bootstrap();
			return;
		}

		if(isServiceDiscovery) {
			discoveryServiceBootstraps();
			return;
		}

	}
}
