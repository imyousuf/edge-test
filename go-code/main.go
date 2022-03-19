package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", helloWorld)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + os.Getenv(("PORT")),
	}
	if serverListenErr := server.ListenAndServe(); serverListenErr != nil {
		fmt.Println(serverListenErr)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	fmt.Fprint(w, "Hello World from GoLang!")
}
