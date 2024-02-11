package main

import (
	"MSHUGO/proxy/internal/controller/auth"
	"MSHUGO/proxy/internal/controller/geo"
	"MSHUGO/proxy/internal/controller/user"
	"MSHUGO/proxy/internal/grpc/grpcclient"
	"MSHUGO/proxy/internal/router"
	"log"
	"net/http"
)

func main() {
	gcl := grpcclient.NewClientGeo()
	acl := grpcclient.NewClientAuth()
	ucl := grpcclient.NewClientUser()
	hg := geo.NewHandGeo(gcl)
	ah := auth.NewHandleAuth(acl)
	uh := user.NewHandleUser(ucl)
	r := router.StRout(hg, ah, uh)

	log.Println("proxyq serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}
