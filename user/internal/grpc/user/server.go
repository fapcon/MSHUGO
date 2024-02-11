package user

import (
	"MSHUGO/user/internal/models"
	"MSHUGO/user/internal/service"
	"context"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
)

type ServicerUser interface {
	CreateUser(email, password string) (string, error)
	CheckUser(email, password string) error
	ProfileUser(email string) (*models.UserDTO, error)
	ListUsers() (*[]models.UserDTO, error)
}

type ServiceUser struct {
	userpr.UnimplementedUserServiceServer
	us *service.UserService
}

func NewServiceUser(usservice *service.UserService) *ServiceUser {
	return &ServiceUser{us: usservice}
}

func (s *ServiceUser) CreateUser(ctx context.Context, req *userpr.CreateRequest) (*userpr.CreateResponse, error) {
	message, err := s.us.Create(req.Email, req.Hashedpassword)
	if err != nil {
		return nil, err
	}
	return &userpr.CreateResponse{Message: message}, nil
}

func (s *ServiceUser) CheckUser(ctx context.Context, req *userpr.CheckRequest) (*userpr.CheckResponse, error) {
	err := s.us.Check(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &userpr.CheckResponse{}, nil
}

func (s *ServiceUser) ProfileUser(ctx context.Context, req *userpr.ProfileRequest) (*userpr.ProfileResponse, error) {
	user, err := s.us.Profile(req.Email)
	if err != nil {
		return nil, err
	}
	p := &userpr.User{Id: user.Id, Email: user.Email}

	return &userpr.ProfileResponse{User: p}, nil
}
func (s *ServiceUser) ListUsers(ctx context.Context, req *userpr.ListRequest) (*userpr.ListResponse, error) {
	users, err := s.us.List()
	if err != nil {
		return nil, err
	}
	var grpcUsers []*userpr.User
	for _, user := range users {
		grpcUser := &userpr.User{
			Id:    user.Id,
			Email: user.Email,
		}
		grpcUsers = append(grpcUsers, grpcUser)
	}
	return &userpr.ListResponse{Users: grpcUsers}, nil

}
