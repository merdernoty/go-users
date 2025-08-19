package transport

import (
	"context"
	pb "github.com/merdernoty/anime-proto/anime"
	"github.com/merdernoty/microservices-planner/anime-search/internal/anime/domain"
)

type AnimeGRPC struct {
	pb.UnimplementedAnimeServiceServer
	as domain.AnimeService
}

func NewAnimeGRPC(as domain.AnimeService) *AnimeGRPC {
	return &AnimeGRPC{as: as}
}

func (g *AnimeGRPC) SearchAnime(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	results, err := g.as.Search(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	var out pb.SearchResponse
	for _, a := range results {
		out.Results = append(out.Results, &pb.AnimeItem{
			Id: int32(a.ID), Title: a.Title, Synopsis: a.Synopsis,
		})
	}
	return &out, nil
}

func (g *AnimeGRPC) GetAnimeById(ctx context.Context, req *pb.GetByIdRequest) (*pb.AnimeResponse, error) {
	res, err := g.as.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &pb.AnimeResponse{}, nil
	}
	return &pb.AnimeResponse{Anime: &pb.AnimeItem{Id: int32(res.ID), Title: res.Title, Synopsis: res.Synopsis}}, nil
}

func (g *AnimeGRPC) ListAnime(ctx context.Context, req *pb.ListRequest) (*pb.SearchResponse, error) {
	results, err := g.as.List(ctx)
	if err != nil {
		return nil, err
	}

	var out pb.SearchResponse
	for _, a := range results {
		out.Results = append(out.Results, &pb.AnimeItem{
			Id: int32(a.ID), Title: a.Title, Synopsis: a.Synopsis,
		})
	}
	return &out, nil
}
