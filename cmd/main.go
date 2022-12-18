package main

import (
	"go_grpc/cmd/config"
	"go_grpc/cmd/service"
	productPb "go_grpc/pb/product"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listend %v", err.Error())
	}

	db := config.ConnectDB()

	grpcServer := grpc.NewServer()
	productService := service.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
