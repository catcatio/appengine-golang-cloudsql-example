package app

import (
	"fmt"
	"github.com/catcatio/shio/app/database"
	"github.com/catcatio/shio/migrations"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func Bootstrap() {
	var (
		r = mux.NewRouter()
	)

	db := database.MakeCloudSQLProxyDBConnection()
	migrations.Migrate(db)
	r.Methods("GET").PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("{\"message\": \"HI\"}"))
	})

	srv := http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	fmt.Println("Start server on :8080")
	log.Fatal(srv.ListenAndServe())
}
