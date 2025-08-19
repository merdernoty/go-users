package service

import (
	"context"

	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/client"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/domain"
)

type AnimeService struct {
	jikan *client.JikanClient
}

func NewAnimeService(jikan *client.JikanClient) *AnimeService {
	return &AnimeService{jikan: jikan}
}

var _ domain.AnimeService = (*AnimeService)(nil)

func (s *AnimeService) Search(ctx context.Context, query string) ([]*domain.Anime, error) {
	items, err := s.jikan.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Anime, 0, len(items))
	for _, a := range items {
		out = append(out, &domain.Anime{ID: a.ID, Title: a.Title, Synopsis: a.Synopsis})
	}
	return out, nil
}

func (s *AnimeService) GetById(ctx context.Context, id int) (*domain.Anime, error) {
	item, err := s.jikan.GetByID(ctx, id)
	if err != nil || item == nil {
		return nil, err
	}
	return &domain.Anime{ID: item.ID, Title: item.Title, Synopsis: item.Synopsis}, nil
}

func (s *AnimeService) List(ctx context.Context) ([]*domain.Anime, error) {
	items, err := s.jikan.Search(ctx, "") 
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Anime, 0, len(items))
	for _, a := range items {
		out = append(out, &domain.Anime{ID: a.ID, Title: a.Title, Synopsis: a.Synopsis})
	}
	return out, nil
}
