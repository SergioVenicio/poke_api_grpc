package grpc

import (
	"context"
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
	response, err := s.GetPokemonUseCase.Do(name)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	var moves []*pb.Move
	for _, m := range response.Moves {
		var details []*pb.VersionGroupDetail
		for _, d := range m.VersionGroupDetails {
			details = append(details, &pb.VersionGroupDetail{
				LevelLearnedAt: d.LevelLearnedAt,
				VersionGroup: &pb.GroupDetail{
					Name: d.VersionGroup.Name,
					Url:  d.VersionGroup.Url,
				},
			})
		}
		moves = append(moves, &pb.Move{
			Move: &pb.MoveDetail{
				Name: m.MoveDetails.Name,
				Url:  m.MoveDetails.Url,
			},
			VersionGroupDetails: details,
		})
	}

	var abilities []*pb.AbilityResponse
	for _, a := range response.Abilities {
		abilities = append(abilities, &pb.AbilityResponse{
			Ability: &pb.Ability{
				Name: a.AbilityDetail.Name,
				Url:  a.AbilityDetail.Url,
			},
			IsHidden: a.IsHidden,
			Slot:     a.Slot,
		})
	}

	var forms []*pb.Form
	for _, f := range response.Forms {
		forms = append(forms, &pb.Form{
			Name: f.Name,
			Url:  f.Url,
		})
	}

	var gameIndices []*pb.GameIndex
	for _, g := range response.GameIndices {
		gameIndices = append(gameIndices, &pb.GameIndex{
			GameIndex: g.GameIndex,
			Version: &pb.Version{
				Name: g.Version.Name,
				Url:  g.Version.Url,
			},
		})
	}

	var heldItems []*pb.HeldItem
	for _, h := range response.HeldItems {
		var versions []*pb.VersionDetail
		for _, v := range h.HeldDetail {
			versions = append(versions, &pb.VersionDetail{
				Rarity: v.Rarity,
				Version: &pb.Version{
					Name: v.Version.Name,
					Url:  v.Version.Url,
				},
			})
		}
		heldItems = append(heldItems, &pb.HeldItem{
			Item: &pb.Item{
				Name: h.Item.Name,
				Url:  h.Item.Url,
			},
			VersionDetail: versions,
		})
	}

	return &pb.GetPokemonResponse{
		Id:             response.Id,
		Name:           response.Name,
		BaseExperience: response.BaseExperience,
		Height:         response.Height,
		Abilities:      abilities,
		Forms:          forms,
		GameIndices:    gameIndices,
		HeldItems:      heldItems,
		Moves:          moves,
	}, nil
}

func NewPokeApiServer(listUseCase *usecases.ListPokemonsUseCase, getUseCase *usecases.GetPokemonUseCase) *PokeApiServer {
	return &PokeApiServer{
		ListUseCase:       listUseCase,
		GetPokemonUseCase: getUseCase,
	}
}
