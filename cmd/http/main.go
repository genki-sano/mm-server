package main

import (
	"os"

	"github.com/genki-sano/mm-server/cmd/http/route"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := route.Route().Run(":" + port); err != nil {
		panic(err)
	}
}
