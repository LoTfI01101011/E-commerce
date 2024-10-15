package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LoTfI01101011/E-commerce/User_service/api/gRPC"
	pb "github.com/LoTfI01101011/E-commerce/User_service/api/gRPC/proto"
	"github.com/LoTfI01101011/E-commerce/User_service/internal"

	"google.golang.org/grpc"
)

func main() {
	internal.DbConnection()
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("fatal to listen: %w", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &gRPC.Server{})

	fmt.Println("We'r starting the server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed runing the server")
	}

}
