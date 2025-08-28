import { envValidate } from '@/core/utils'
import { ConfigModuleOptions } from '@nestjs/config'
import { Type } from 'class-transformer'
import { IsInt } from 'class-validator'

export class EnvironmentServiceDiscoveryVariables {
    @Type(() => Number)
    @IsInt()
    PORT: number
}

export const EnvServiceDiscoveryConfigOptions: ConfigModuleOptions = {
    validate: envValidate(EnvironmentServiceDiscoveryVariables),
    isGlobal: true,
}