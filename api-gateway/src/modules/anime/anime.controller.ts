import { AnimeService } from '@/generated/anime/anime-search'
import { ServiceDiscovery } from '@/service-discovery'
import { ServiceDiscoveryClient } from '@/service-discovery/client'
import { Controller, Get, Param, ParseIntPipe, Query, UsePipes, ValidationPipe } from '@nestjs/common'

@UsePipes(new ValidationPipe({ whitelist: true }))
@Controller('anime')
export class AnimeController {
	constructor(private serviceDiscoveryClient: ServiceDiscoveryClient ) {}

	private async animeService() {
		const result = await this.serviceDiscoveryClient.getClient('ANIME_CLIENT')
		if(!result.isSuccess || !result.client) {
			throw new Error('Anime service is unavailable')
		}

		return result.client.getService<AnimeService>('AnimeService')
	}

	@Get(':id')
	public async getOne(@Param('id', ParseIntPipe) id: number) {
		const animeService = await this.animeService()
		const response = await animeService.GetAnimeById({ id })
		return response
	}

	@Get()
	public async getList() {
		const animeService = await this.animeService()
		const response = await animeService.ListAnime({})
		return response
	}

	@Get('search')
	public async search(@Query('query') query: string) {
		const animeService = await this.animeService()
		const response = await animeService.SearchAnime({ query })
		return response
	}
}
