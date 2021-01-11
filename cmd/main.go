package main

import (
	"log"
	"net"

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
	address := ":8080"

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

	userRepository := datastore.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)

	userUseCase := usecase.NewUserUseCase(userRepository, transaction, userService)
	userHandler := handler.NewUserHandler(userUseCase)

	lis, err := net.Listen("tcp", address)
	server := grpc.NewServer() // ここでInterceptorとか入れる
	rpc.RegisterUserServiceServer(server, userHandler)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Print("aaaaaa")
		log.Fatalf("failed to serve: %s", err)
	}

}

// func main() {
// 	// port := 8080

// 	db, err := datastore.NewMysqlDB(config.DSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer func() {
// 		err := db.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Sample")
//     })
//     http.ListenAndServe(":8080", nil)
// }
