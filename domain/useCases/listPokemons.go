package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ListPokemonsUseCase struct {
	baseUrl string
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
	url := fmt.Sprintf("%s/?limit=%d&offset=%d", l.baseUrl, limit, offset)
	log.Println(fmt.Sprintf("Fetching list of pokemons (%s)", url))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var response ListResponse
	json.NewDecoder(res.Body).Decode(&response)
	return &response, nil
}

func NewListUseCase() *ListPokemonsUseCase {
	return &ListPokemonsUseCase{
		baseUrl: os.Getenv("POKE_API_BASE_UR"),
	}
}
