package grpc

import (
	"context"
	grpcMapper "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/mappers/pokemon/grpc"
	usecases "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/useCases"
	pb "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/grpc/proto"
	"log"
)

type PokeApiServer struct {
	pb.UnimplementedPokeAPIServer
	ListUseCase       *usecases.ListPokemonsUseCase
	GetPokemonUseCase *usecases.GetPokemonUseCase
}

func (s *PokeApiServer) ListPokemons(ctx context.Context, request *pb.ListPokemonRequest) (*pb.ListPokemonsResponse, error) {
	limit := request.GetLimit()
	if limit == 0 {
		limit = 20
	}
	offset := request.GetOffset()
	response, err := s.ListUseCase.Do(limit, offset)
	if err != nil {
		return nil, err
	}

	var results []*pb.ListPokemonData
	for _, r := range response.Results {
		results = append(results, &pb.ListPokemonData{
			Name: r.Name,
			Url:  r.Url,
		})
	}

	return &pb.ListPokemonsResponse{
		Results:  results,
		Count:    int64(response.Count),
		Next:     response.Next,
		Previous: response.Previous,
	}, nil
}

func (s *PokeApiServer) GetPokemon(ctx context.Context, request *pb.GetPokemonRequest) (*pb.GetPokemonResponse, error) {
	name := request.GetName()
	mapper := grpcMapper.NewPokemonGRPCMapper()
	response, err := s.GetPokemonUseCase.Do(name)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return mapper.FromDomain(response), nil
}

func NewPokeApiServer(listUseCase *usecases.ListPokemonsUseCase, getUseCase *usecases.GetPokemonUseCase) *PokeApiServer {
	return &PokeApiServer{
		ListUseCase:       listUseCase,
		GetPokemonUseCase: getUseCase,
	}
}
