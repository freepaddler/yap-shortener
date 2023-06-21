package main

import (
	"fmt"
	"net/http"

	"github.com/freepaddler/yap-shortener/internal/app/config"
	"github.com/freepaddler/yap-shortener/internal/app/handlers"
	"github.com/freepaddler/yap-shortener/internal/app/router"
	"github.com/freepaddler/yap-shortener/internal/app/store"
)

func main() {
	fmt.Println("Starting server...")
	conf := config.NewConfig()
	s := store.NewMemStore()
	h := handlers.NewHTTPHandler(s, conf.ServerURL)
	r := router.NewHTTPRouter(h)

	if err := http.ListenAndServe(conf.ServerAddress, r); err != nil {
		panic(err)
	}

	fmt.Println("Stopping server...")

}
