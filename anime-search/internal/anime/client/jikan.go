package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.jikan.moe/v4"

type JikanClient struct {
	http *http.Client
}

func NewJikanClient() *JikanClient {
	return &JikanClient{http: &http.Client{Timeout: 10 * time.Second}}
}

type AnimeItem struct {
	ID       int    `json:"mal_id"`
	Title    string `json:"title"`
	Synopsis string `json:"synopsis"`
}

type searchResponse struct {
	Data []AnimeItem `json:"data"`
}

type getByIDResponse struct {
	Data AnimeItem `json:"data"`
}

func (c *JikanClient) Search(ctx context.Context, q string) ([]AnimeItem, error) {
	u := fmt.Sprintf("%s/anime?q=%s", baseURL, url.QueryEscape(q))
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("jikan rate limited: %s", resp.Status)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("jikan search error: %s", resp.Status)
	}

	var out searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out.Data, nil
}

func (c *JikanClient) GetByID(ctx context.Context, id int) (*AnimeItem, error) {
	u := fmt.Sprintf("%s/anime/%d", baseURL, id)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("jikan rate limited: %s", resp.Status)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("jikan getById error: %s", resp.Status)
	}

	var out getByIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out.Data, nil
}
