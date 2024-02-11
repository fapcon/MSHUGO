package grpcclient

import (
	"context"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
	"google.golang.org/grpc"
	"log"
)

type ClientGeo struct{}

func NewClientGeo() *ClientGeo {
	return &ClientGeo{}
}

func (c *ClientGeo) CallSearchAddress(ctx context.Context, req *geopr.SearchRequest) (*geopr.SearchResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := geopr.NewGeoServiceClient(conn)

	res, err := client.Search(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове gRPC: %v", err)
		return nil, err
	}

	return res, nil
}

func (c *ClientGeo) CallGeocodeAddress(ctx context.Context, req *geopr.GeocodeRequest) (*geopr.GeocodeResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := geopr.NewGeoServiceClient(conn)

	res, err := client.Geocode(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове gRPC: %v", err)
		return nil, err
	}

	return res, nil
}
