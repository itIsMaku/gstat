package httpserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gstat/internal/configuration"
	"gstat/internal/database"
	"net/http"
)

type Status struct {
	Alive bool `json:"alive"`
}

func Start(httpConfig configuration.Http, db *sql.DB) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(Status{Alive: true})
	})

	http.HandleFunc("/history", func(writer http.ResponseWriter, request *http.Request) {
		results, err := database.GetResultsBefore(db)
		if err != nil {
			http.Error(writer, "Error retrieving history", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(results)
	})

	fmt.Println("Starting HTTP server on port", httpConfig.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpConfig.Port), nil)
}
