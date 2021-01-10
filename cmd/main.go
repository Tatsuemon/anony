package main

import (
    "fmt"
	"net/http"
	
	"github.com/Tatsuemon/shortURL/infrastructure/datastore"
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
	}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Sample")
    })
    http.ListenAndServe(":8080", nil)
}
