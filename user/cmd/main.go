package main

import (
	"MSHUGO/user/internal/grpc/user"
	"MSHUGO/user/internal/repository"
	"MSHUGO/user/internal/service"
	userpr "github.com/fapcon/MSHUGOprotos/protos/user/gen"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	// Подключение к базе данных
	time.Sleep(2 * time.Second)
	dbHost := "db"
	dbPort := "5432"
	dbUser := "userpostgres"
	dbPassword := "password"
	dbName := "userserv"
	sslmode := "disable"

	connectionString := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + sslmode

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	if err != nil {
		log.Fatalf("ping:%v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        hashepassword VARCHAR(255) NOT NULL
    )`)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)

	serviceUser := user.NewServiceUser(userService)

	// Создание gRPC сервера
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", "50053", err)
	}
	grpcServer := grpc.NewServer()

	// Регистрация ServiceUser в gRPC сервере
	userpr.RegisterUserServiceServer(grpcServer, serviceUser)

	// Запуск gRPC сервера
	log.Print("Starting gRPC server user...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
