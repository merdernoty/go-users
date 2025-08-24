import { GrpcOptions, Transport } from '@nestjs/microservices'
import { join } from 'path'

export const userGrpcConfig: GrpcOptions = {
	transport: Transport.GRPC,
	options: {
		package: 'user',
		protoPath: join(__dirname, '../../../anime-proto/user/user.proto')
	}
}
