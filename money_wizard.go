package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	apiv1 "github.com/nakonechniyd/money_wizard/api/v1"
	"github.com/nakonechniyd/money_wizard/strings"
)

func main() {
	host := strings.StrOrStr(os.Getenv("MW_HOST"), "localhost")
	port := strings.StrOrStr(os.Getenv("MW_PORT"), "8080")

	router := mux.NewRouter().StrictSlash(true)
	apiv1.AddRouts(router)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Handler:      router,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	fmt.Printf("Running on http://%s:%s\n", host, port)
	log.Fatal(srv.ListenAndServe())
}
