package grpcclient

import (
	"context"
	authpr "github.com/fapcon/MSHUGOprotos/protos/auth/gen"
	"google.golang.org/grpc"
	"log"
)

type ClientAuth struct{}

func NewClientAuth() *ClientAuth {
	return &ClientAuth{}
}

func (c *ClientAuth) CallRegister(ctx context.Context, req *authpr.RegisterRequest) (*authpr.RegisterResponse, error) {
	conn, err := grpc.Dial("auth:44972", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := authpr.NewAuthServiceClient(conn)

	res, err := client.Register(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}

	log.Printf("Ответ от сервера auth: %s", res.Message)
	return res, nil
}

func (c *ClientAuth) CallLogin(ctx context.Context, req *authpr.LoginRequest) (*authpr.LoginResponse, error) {
	conn, err := grpc.Dial("auth:44972", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := authpr.NewAuthServiceClient(conn)

	res, err := client.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}

	log.Printf("Ответ от сервера auth: %s", res.Token)
	return res, nil
}
