package router

import (
	"MSHUGO/proxy/internal/controller/auth"
	"MSHUGO/proxy/internal/controller/geo"
	"MSHUGO/proxy/internal/controller/user"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func StRout(geohand *geo.HandleGeo, authhand *auth.HandleAuth, userhand *user.HandleUser) *chi.Mux {

	r := chi.NewRouter()

	rp := NewReverseProxy("hugo", "1313")
	r.Use(rp.ReverseProxy)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Post("/api/login", authhand.Login)
	r.Post("/api/register", authhand.Register)

	r.Group(func(r chi.Router) {
		r.Use(TokenValidationMiddleware)
		r.Get("/api/profile", userhand.ProfileUser)
		r.Get("/api/list", userhand.ListUsers)
		r.Post("/api/address/search", geohand.SearchHandle)
		r.Post("/api/address/geocode", geohand.GeocodeHandle)
	})

	return r
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {

	target, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
	proxy := httputil.NewSingleHostReverseProxy(target)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/swagger") && !strings.HasPrefix(r.URL.Path, "/api") {
			proxy.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secretkey"), nil
		})

		if err != nil {
			http.Error(w, "invalid token", http.StatusForbidden)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			ctx := context.WithValue(r.Context(), "claims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "invalid token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
