package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	usecases "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/domain/useCases"
	grpcInfra "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/grpc"
	pb "github/sergiovenicio/poke_apigithub/sergiovenicio/poke_api/infra/grpc/proto"
)

func main() {
	godotenv.Load()

	listUseCase := usecases.NewListUseCase()
	getUseCase := usecases.NewGetPokemonUseCase()
	lis, _ := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	fmt.Println(lis.Addr())

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	pb.RegisterPokeAPIServer(server, grpcInfra.NewPokeApiServer(listUseCase, getUseCase))
	server.Serve(lis)
}
