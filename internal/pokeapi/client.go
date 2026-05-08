package pokeapi

import (
	"net/http"
	"time"

	"github.com/Lando-Iraola/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	pokecache := pokecache.NewCache(time.Second * 15)

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache,
	}
}
