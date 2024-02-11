package grpcclients

import (
	"context"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"google.golang.org/grpc"
	"log"
)

type ClientUser struct{}

func (c *ClientUser) CallCreateUser(ctx context.Context, req *userpr.CreateRequest) (*userpr.CreateResponse, error) {
	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := userpr.NewUserServiceClient(conn)

	res, err := client.Create(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc")
	}
	return res, nil
}

func (c *ClientUser) CallCheckUser(ctx context.Context, req *userpr.CheckRequest) (*userpr.CheckResponse, error) {
	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := userpr.NewUserServiceClient(conn)

	res, err := client.Check(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc")
	}
	return res, nil
}

func (c *ClientUser) CallProfileUser(ctx context.Context, req *userpr.ProfileRequest) (*userpr.ProfileResponse, error) {
	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := userpr.NewUserServiceClient(conn)

	res, err := client.Profile(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc")
	}
	return res, nil
}
