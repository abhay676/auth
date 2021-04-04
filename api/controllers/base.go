package controllers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Server struct {
	DB     *mongo.Database
	Router *mux.Router
}

func (s *Server) Initialize(uri, database string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("DB error %v", err)
	}
	s.DB = client.Database(database)
	fmt.Println("DB Connected")
	s.Router = mux.NewRouter()
	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Printf("Listening on Port %v", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}