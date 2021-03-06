package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Tatsuemon/anony/infrastructure/middleware"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/web/handler"
	"github.com/Tatsuemon/anony/rpc"
	"github.com/Tatsuemon/anony/usecase"

	"github.com/Tatsuemon/anony/config"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("API_PORT")

	db, err := datastore.NewMysqlDB(config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	transaction := datastore.NewTransaction(db.DB)

	// User
	userRepository := datastore.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)

	userUseCase := usecase.NewUserUseCase(userRepository, transaction, userService)
	userHandler := handler.NewUserHandler(userUseCase)

	// AnonyURL
	anonyURLRepository := datastore.NewAnonyURLRepository(db.DB)
	anonyURLService := service.NewAnonyURLService(anonyURLRepository)

	userAnonyURLAccessor := datastore.NewUserAnonyURLAccessor(db.DB)

	anonyURLUseCase := usecase.NewAnonyURLUseCase(anonyURLRepository, transaction, anonyURLService)

	anonyWithUserUseCase := usecase.NewAnonyURLWithUserUseCase(userAnonyURLAccessor, transaction)

	anonayURLHandler := handler.NewAnonyURLHandler(anonyURLUseCase, anonyWithUserUseCase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(middleware.UnaryServerInterceptor(middleware.JWTAuth(userService))),
	) // ここでInterceptorとか入れる

	rpc.RegisterUserServiceServer(server, userHandler)
	rpc.RegisterAnonyServiceServer(server, anonayURLHandler)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Print("aaaaaa")
		log.Fatalf("failed to serve: %s", err)
	}

}
