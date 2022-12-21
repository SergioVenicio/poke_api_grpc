package usecases

import (
	"encoding/json"
	"fmt"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/redis"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/settings"
	"log"
	"net/http"
	"os"
	"time"
)

type ListPokemonsUseCase struct {
	baseUrl  string
	redis    *redis.Redis
	cacheTTL int
}

type pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ListResponse struct {
	Results  []pokemon `json:"results"`
	Count    int       `json:"count"`
	Previous string    `json:"previous"`
	Next     string    `json:"next"`
}

func (l *ListPokemonsUseCase) Do(limit int64, offset int64) (*ListResponse, error) {
	var response ListResponse
	url := fmt.Sprintf("%s/?limit=%d&offset=%d", l.baseUrl, limit, offset)
	cached := l.redis.Get(url)
	if cached != "" {
		log.Println("Using cache for", url, "...")
		err := json.Unmarshal([]byte(cached), &response)
		if err == nil {
			return &response, nil
		}
	}

	log.Println(fmt.Sprintf("Fetching list of pokemons (%s)", url))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	json.NewDecoder(res.Body).Decode(&response)
	l.redis.Set(url, response, time.Hour*4)
	return &response, nil
}

func NewListUseCase() *ListPokemonsUseCase {
	return &ListPokemonsUseCase{
		baseUrl:  os.Getenv("POKE_API_BASE_URI"),
		redis:    redis.NewRedis(),
		cacheTTL: settings.LoadCacheTTL(),
	}
}
