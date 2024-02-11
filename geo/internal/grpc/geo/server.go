package geo

import (
	"MSHUGO/geo/internal/grpc/grpcclients"
	"MSHUGO/geo/internal/service"
	"context"
	"fmt"
	geopr "github.com/fapcon/MSHUGOprotos/protos/geo/gen"
)

type Geo interface {
	GeoSearch(input string) ([]byte, error)
	GeoCode(lat, lng string) ([]byte, error)
}

type ServerGeo struct {
	geopr.UnimplementedGeoServiceServer
	gcl *grpcclients.ClientAuth
	geo service.GeoService
}

func (s *ServerGeo) SearchAddress(context context.Context, req *geopr.SearchRequest) (*geopr.SearchResponse, error) {

	address, err := s.geo.GeoSearch(req.Input)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}

	return &geopr.SearchResponse{Data: address}, nil
}

func (s *ServerGeo) GeocodeAddress(context context.Context, req *geopr.GeocodeRequest) (*geopr.GeocodeResponse, error) {
	address, err := s.geo.GeoCode(req.Lat, req.Lon)
	if err != nil {
		return nil, fmt.Errorf("err service:%v", err)
	}
	return &geopr.GeocodeResponse{Data: address}, nil
}
