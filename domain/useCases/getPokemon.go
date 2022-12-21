package usecases

import (
	"encoding/json"
	"fmt"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/entities"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/redis"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/settings"
	"log"
	"net/http"
	"os"
	"time"
)

type GetPokemonUseCase struct {
	baseUrl  string
	redis    *redis.Redis
	cacheTTL int
}

func (u *GetPokemonUseCase) Do(name string) (*entities.Pokemon, error) {
	var response entities.Pokemon
	url := fmt.Sprintf("%s/%s", u.baseUrl, name)
	cached := u.redis.Get(url)
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
		log.Fatalf(err.Error())
		return nil, err
	}

	json.NewDecoder(res.Body).Decode(&response)
	u.redis.Set(url, response, time.Hour*4)
	return &response, nil
}

func NewGetPokemonUseCase() *GetPokemonUseCase {
	return &GetPokemonUseCase{
		baseUrl:  os.Getenv("POKE_API_BASE_URI"),
		redis:    redis.NewRedis(),
		cacheTTL: settings.LoadCacheTTL(),
	}
}
