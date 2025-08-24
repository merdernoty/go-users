import cluster from 'cluster'
import { cpus } from 'os'


// TODO: Проблема в Cluster mode, 
// TODO: в том что не будет общий Discovery Service -> либо в отдельный воркер или в хранить записи 
// TODO: в Redis/memchached если возможно
export const clusterModeBootstrap = (bootstrap: () => Promise<void> | void) => {
	if (cluster.isPrimary) {
		const numCPUs = cpus().length

		for (let i = 0; i < numCPUs; i++) {
			cluster.fork()
		}

		cluster.on('exit', (worker, code, signal) => {
			// LOGGER
			cluster.fork()
		})
	} else {
		bootstrap()
	}
}
