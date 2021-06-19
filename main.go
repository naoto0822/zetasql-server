package main

/*
#include <stdlib.h>
#include "parse_query.h"
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"unsafe"
)

func main() {
	fmt.Println("zetasql-ast-server v0.0.1")

	http.HandleFunc("/valid", validHandler)
	http.HandleFunc("/ast", astHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func validHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cs := C.CString(string(b))
	defer C.free(unsafe.Pointer(cs))

	formatResult := C.isValidStatement(cs)
	defer C.free(unsafe.Pointer(formatResult))

	w.Write([]byte(C.GoString(formatResult)))
}

func astHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cs := C.CString(string(b))
	defer C.free(unsafe.Pointer(cs))

	formatResult := C.parseStatement(cs)
	defer C.free(unsafe.Pointer(formatResult))

	w.Write([]byte(C.GoString(formatResult)))
}
