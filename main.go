package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/phaseharry/learn-go/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load() // loads env variables from .env file

	// getting "PORT" env variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment!")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment!")
	}
	fmt.Println("port: " + port + "\ndbURL: " + dbUrl)

	// Go standard library has built in SQL package
	dbConnection, dbConnectionErr := sql.Open("postgres", dbUrl)
	if dbConnectionErr != nil {
		log.Fatal("DB_URL is not found in the environment!")
	}

	apiCfg := apiConfig{
		DB: database.New(dbConnection),
	}

	router := chi.NewRouter()

	// cors config so clients can make requests to server from browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// creating a secondary router for "/ready" path and mounting it to the
	// main router's "/v1" path
	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)
	/*
		blocks and starts handling HTTP requests. If there's any error
		then it will be unblocked and return that error
	*/
	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal(serverError)
	}
}

// initializing a go module
// go mod init ${name-of-module}

// grabbing a go package to use
// go get github.com/joho/godotenv

// loading dependencies into its own directory
/*
	go mod vendor
	- first call creates the vendor directory with a list of dependencies.
	- second call imports all of the listed dependencies
*/
