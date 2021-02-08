package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Tatsuemon/anony/config"
	"github.com/Tatsuemon/anony/domain/service"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
	"github.com/Tatsuemon/anony/infrastructure/web/handler"
	"github.com/Tatsuemon/anony/usecase"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("HTTP_PORT")
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
	anonyURLRepository := datastore.NewAnonyURLRepository(db.DB)
	anonyURLService := service.NewAnonyURLService(anonyURLRepository)
	anonyURLUseCase := usecase.NewAnonyURLUseCase(anonyURLRepository, transaction, anonyURLService)

	mux := mux.NewRouter()
	catchAllHandler := handler.NewHttpHandler(anonyURLUseCase)
	mux.PathPrefix("/").Handler(catchAllHandler)
	fmt.Printf("Server running at http://loacalhost:%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatal(err)
	}
}
