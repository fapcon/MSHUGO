package user

import (
	"context"
	"encoding/json"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"log"
	"net/http"
	"proxy/internal/grpc/grpcclient"
	"proxy/internal/models"
)

type HandleUser struct {
	clientuser *grpcclient.ClientUser
}

func NewHandleUser(clUser *grpcclient.ClientUser) *HandleUser {
	return &HandleUser{clUser}
}

func (h *HandleUser) Profile(w http.ResponseWriter, r *http.Request) {
	email := "qwer"
	req := &userpr.ProfileRequest{Email: email}

	res, err := h.clientuser.CallProfile(context.Background(), req)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	user := &models.User{
		Id:    res.User.Id,
		Email: res.User.Email,
	}
	jsData, err := json.Marshal(user)
	if err != nil {
		log.Println("err:", err)
	}
	w.Write(jsData)
}

func (h *HandleUser) List(w http.ResponseWriter, r *http.Request) {

	req := &userpr.ListRequest{}
	res, err := h.clientuser.CallList(context.Background(), req)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	var users []models.User
	for _, v := range res.Users {
		user := models.User{
			Id:    v.Id,
			Email: v.Email,
		}
		users = append(users, user)
	}
	jsData, err := json.Marshal(users)
	if err != nil {
		log.Println("err :", err)
	}
	w.Write(jsData)
}
