package service

import (
	"fmt"

	"log"
	"user/internal/models"
	"user/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) Create(email, hashepassword string) (string, error) {
	err := u.Repo.Create(email, hashepassword)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	return fmt.Sprint("user created successfully"), nil
}

func (u *UserService) Check(email, password string) error {
	err := u.Repo.Check(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Profile(email string) (*models.UserDTO, error) {
	user, err := u.Repo.Profile(email)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserService) List() ([]models.UserDTO, error) {
	users, err := u.Repo.List()
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return users, nil
}
