package grpc

import (
	"github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/entities"
	pb "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/grpc/proto"
)

type PokemonGRPCMapper struct{}

func (m *PokemonGRPCMapper) FromDomain(pokemon *entities.Pokemon) *pb.GetPokemonResponse {
	var moves []*pb.Move
	for _, m := range pokemon.Moves {
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
	for _, a := range pokemon.Abilities {
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
	for _, f := range pokemon.Forms {
		forms = append(forms, &pb.Form{
			Name: f.Name,
			Url:  f.Url,
		})
	}

	var gameIndices []*pb.GameIndex
	for _, g := range pokemon.GameIndices {
		gameIndices = append(gameIndices, &pb.GameIndex{
			GameIndex: g.GameIndex,
			Version: &pb.Version{
				Name: g.Version.Name,
				Url:  g.Version.Url,
			},
		})
	}
	var heldItems []*pb.HeldItem
	for _, h := range pokemon.HeldItems {
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
		Id:             pokemon.Id,
		Name:           pokemon.Name,
		BaseExperience: pokemon.BaseExperience,
		Height:         pokemon.Height,
		Abilities:      abilities,
		Forms:          forms,
		GameIndices:    gameIndices,
		HeldItems:      heldItems,
		Moves:          moves,
	}
}

func NewPokemonGRPCMapper() *PokemonGRPCMapper {
	return &PokemonGRPCMapper{}
}
