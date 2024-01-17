package main

import (
	"database/sql"
	"net"

	"github.com/VituSuperMEg/go-grpc-fullcycle/internal/database"
	"github.com/VituSuperMEg/go-grpc-fullcycle/internal/pb"
	"github.com/VituSuperMEg/go-grpc-fullcycle/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Registrar o serviço de reflection
	reflection.Register(grpcServer)

	// Abrir uma porta/conexão tcp
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
