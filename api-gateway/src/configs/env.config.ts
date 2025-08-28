import { envValidate } from '@/core/utils'
import { ConfigModuleOptions } from '@nestjs/config'
import { Type } from 'class-transformer'
import { IsInt } from 'class-validator'

export class EnvironmentVariables {
	@Type(() => Number)
    @IsInt()
    PORT: number
}

export const EnvConfigOptions: ConfigModuleOptions = {
	validate: envValidate(EnvironmentVariables),
	isGlobal: true,
}