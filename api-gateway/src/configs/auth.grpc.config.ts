import { GrpcOptions, Transport } from '@nestjs/microservices'
import { join } from 'path'

export const authGrpcConfig: GrpcOptions = {
	transport: Transport.GRPC,
	options: {
		package: 'auth',
		protoPath: join(__dirname, '../../../anime-proto/auth/auth.proto')
	}
}
