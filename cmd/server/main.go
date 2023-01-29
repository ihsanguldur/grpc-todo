package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc-todo/pkg/api/models"
	"grpc-todo/pkg/api/services"
	"grpc-todo/pkg/app"
	"grpc-todo/pkg/pb"
	"grpc-todo/pkg/repository"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var err error

	if err = godotenv.Load("../../configs/.env"); err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := setupDatabase(dsn)
	if err != nil {
		return err
	}
	fmt.Println("server[todo]-database connection established successfully.")

	if err = db.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}
	fmt.Println("server[todo]-migration completed.")

	lis, err := setupNet(os.Getenv("PORT"))
	if err != nil {
		return err
	}
	fmt.Println("server[todo]-started on ", os.Getenv("PORT"))

	storage := repository.NewStorage(db)
	todoService := services.NewTodoService(storage)

	server := app.NewServer(todoService)

	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, server)

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func setupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupNet(port string) (net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return lis, nil
}
