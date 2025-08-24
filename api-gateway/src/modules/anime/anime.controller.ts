import { AnimeService } from '@/generated/anime/anime-search'
import { ServiceDiscovery } from '@/service-discovery'
import { Controller, Get, Param, ParseIntPipe, Query, UsePipes, ValidationPipe } from '@nestjs/common'

@UsePipes(new ValidationPipe({ whitelist: true }))
@Controller('anime')
export class AnimeController {
	constructor(private serviceDiscovery: ServiceDiscovery) {}

	private get animeService() {
		const client = this.serviceDiscovery.getClient('ANIME_CLIENT')
		return client.getService<AnimeService>('AnimeService')
	}

	@Get(':id')
	public async getOne(@Param('id', ParseIntPipe) id: number) {
		const response = await this.animeService.GetAnimeById({ id })
		return response
	}

	@Get()
	public async getList() {
		const response = await this.animeService.ListAnime({})
		return response
	}

	@Get('search')
	public async search(@Query('query') query: string) {
		const response = await this.animeService.SearchAnime({ query })
		return response
	}
}
