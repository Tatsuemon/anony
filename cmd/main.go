package main

import (
    "fmt"
	"net/http"
	"log"
	
	"github.com/Tatsuemon/shortURL/infrastructure/datastore"
	"github.com/Tatsuemon/shortURL/config"
)

func main() {
	// port := 8080

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

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Sample")
    })
    http.ListenAndServe(":8080", nil)
}
