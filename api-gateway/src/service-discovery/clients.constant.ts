import { authGrpcConfig } from '../configs/auth.grpc.config'
import { userGrpcConfig } from '../configs/user.grpc.config'
import { animeGrpcConfig } from '../configs/anime.grpc.config'

export const clientMap = {
	USER_CLIENT: userGrpcConfig,
	AUTH_CLIENT: authGrpcConfig,
	ANIME_CLIENT: animeGrpcConfig
}

export type ClientNameKey = keyof typeof clientMap
