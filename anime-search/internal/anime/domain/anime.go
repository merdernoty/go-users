package domain

import "context"

type Anime struct {
    ID      int
    Title   string
    Synopsis string
}


type AnimeService interface {
	Search(ctx context.Context, query string) ([]*Anime, error)
	GetById(ctx context.Context, id int) (*Anime, error)
	List(ctx context.Context) ([]*Anime, error)
}