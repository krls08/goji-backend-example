package main

import (
	"fmt"
	"os"

	"github.com/carles/repo/server"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Carles string `envconfig:"CARLES"`
}

func main() {
	fmt.Println(os.Getenv("ERIC"))

	server.RunServer()
}
