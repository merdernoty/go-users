import { LoginRequest } from '@/generated/auth/auth'
import { IsEmail, IsString } from 'class-validator'

export class LoginAuthDto implements LoginRequest {
	@IsEmail()
	public email: string

	@IsString()
	public password: string
}
