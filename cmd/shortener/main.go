package main

import (
	"fmt"
	"net/http"

	"github.com/freepaddler/yap-shortener/internal/app/handlers"
	"github.com/freepaddler/yap-shortener/internal/app/router"
	"github.com/freepaddler/yap-shortener/internal/app/store"
)

func main() {
	fmt.Println("Starting server...")
	s := store.NewMemStore()
	h := handlers.NewHTTPHandler(s)
	r := router.NewHTTPRouter(h)

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		panic(err)
	}

	fmt.Println("Stopping server...")

}
