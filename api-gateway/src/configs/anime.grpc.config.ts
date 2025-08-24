import { GrpcOptions, Transport } from '@nestjs/microservices'
import { join } from 'path'

export const animeGrpcConfig: GrpcOptions = {
	transport: Transport.GRPC,
	options: {
		package: 'anime',
		protoPath: join(__dirname, '../../../anime-proto/anime/anime.proto')
	}
}
