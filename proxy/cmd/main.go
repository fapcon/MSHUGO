package main

import (
	"log"
	"net/http"
	"proxy/internal/controller/auth"
	"proxy/internal/controller/geo"
	"proxy/internal/controller/user"
	"proxy/internal/grpc/grpcclient"
	"proxy/internal/router"
)

func main() {
	gcl := grpcclient.NewClientGeo()
	acl := grpcclient.NewClientAuth()
	ucl := grpcclient.NewClientUser()
	hg := geo.NewHandleGeo(gcl)
	ha := auth.NewHandleAuth(acl)
	hu := user.NewHandleUser(ucl)
	r := router.Route(hg, ha, hu)

	log.Println("proxyct serv started on port :8080")
	http.ListenAndServe(":8080", r)
}
