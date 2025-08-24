import { RegisterRequest } from '@/generated/auth/auth'
import { IsEmail, IsString } from 'class-validator'

export class RegisterAuthDto implements RegisterRequest {
	@IsEmail()
	public email: string

    @IsString()
	public password: string

    @IsString()
	public username: string
}
