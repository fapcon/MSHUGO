package main

import (
	"MSHUGO/auth/internal/grpc/auth"
	authpr "github.com/fapcon/MSHUGOprotos/protos/auth/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	authpr.RegisterAuthServiceServer(server, &auth.ServiceAuth{})

	log.Println("Запуск gRPC сервера auth...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
