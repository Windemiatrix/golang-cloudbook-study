package main

import (
	"net/http"

	"github.com/Windemiatrix/golang-cloudbook-study/internal/adapter/rest"
	"github.com/Windemiatrix/golang-cloudbook-study/internal/adapter/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "data",
		FieldMap:          logrus.FieldMap{},
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	logrus.SetLevel(logrus.InfoLevel)

	repo := storage.NewInMemoryRepository()
	handler := rest.NewHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/v1/key/{key}", handler.GetKeyValue).Methods("GET")
	r.HandleFunc("/v1/key/{key}", handler.SetKeyValue).Methods("PUT")
	r.HandleFunc("/v1/key/{key}", handler.DeleteKeyValue).Methods("DELETE")

	logrus.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatalf("Server failed to start: %v", err)
	}
}
