package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load() // loads env variables from .env file

	// getting "PORT" env variable
	port := os.Getenv("PORT")
	fmt.Println("port: " + port)
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
