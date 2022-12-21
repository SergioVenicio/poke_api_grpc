package usecases

import (
	"encoding/json"
	"fmt"
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/entities"
	"log"
	"net/http"
	"os"
)

type GetPokemonUseCase struct {
	baseUrl string
}

func (u *GetPokemonUseCase) Do(name string) (*entities.Pokemon, error) {
	url := fmt.Sprintf("%s/%s", u.baseUrl, name)
	log.Println(fmt.Sprintf("Fetching list of pokemons (%s)", url))
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	var response entities.Pokemon
	json.NewDecoder(res.Body).Decode(&response)
	return &response, nil
}

func NewGetPokemonUseCase() *GetPokemonUseCase {
	return &GetPokemonUseCase{
		baseUrl: os.Getenv("POKE_API_BASE_UR"),
	}
}
