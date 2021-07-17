package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/naoto0822/zetasql-server/pkg"
)

func main() {
	fmt.Printf("zetasql-server:%s\n", pkg.Version)

	http.HandleFunc("/valid", validHandler)
	http.HandleFunc("/parse", parseHandler)
	http.HandleFunc("/format", formatHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func validHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid := IsValidStatement(ctx, string(b))

	w.Write([]byte(isValid))
}

func parseHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	astOrErr := ParseStatement(ctx, string(b))

	w.Write([]byte(astOrErr))
}

func formatHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sqlOrErr := FormatSQL(ctx, string(b))

	w.Write([]byte(sqlOrErr))
}
