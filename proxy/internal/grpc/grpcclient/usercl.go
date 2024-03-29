package grpcclient

import (
	"context"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"google.golang.org/grpc"
	"log"
)

type ClientUser struct{}

func NewClientUser() *ClientUser {
	return &ClientUser{}
}

func (c *ClientUser) CallProfile(ctx context.Context, req *userpr.ProfileRequest) (*userpr.ProfileResponse, error) {

	conn, err := grpc.Dial("user:44971", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := userpr.NewUserServiceClient(conn)
	res, err := client.Profile(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}

func (c *ClientUser) CallList(ctx context.Context, req *userpr.ListRequest) (*userpr.ListResponse, error) {

	conn, err := grpc.Dial("user:44971", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := userpr.NewUserServiceClient(conn)
	res, err := client.List(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}
