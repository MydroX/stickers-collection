package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MydroX/stickers-collection/pkg/db/postgres"
	"github.com/MydroX/stickers-collection/pkg/env"
	"github.com/MydroX/stickers-collection/src/router"
)

func main() {
	env.Load()

	postgres.InitDB()

	r := router.New()

	log.Println(fmt.Sprintf("Server is starting on port %v", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r))
}
