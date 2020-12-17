package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	fr "github.com/thedevsaddam/clean/follower/repository/rest"
	pr "github.com/thedevsaddam/clean/profile/repository/inmemory"
	uHttpDelivery "github.com/thedevsaddam/clean/user/delivery/http"
	ur "github.com/thedevsaddam/clean/user/repository/inmemory"
	"github.com/thedevsaddam/clean/user/usecase"
)

func main() {
	r := chi.NewRouter()
	// middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	u := ur.NewInMemoryUserRepository()
	p := pr.NewInMemoryProfileRepository()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	f := fr.NewrestFollowerRepository(&client)

	uc := usecase.NewUserUsecase(u, p, f)
	uHttpDelivery.NewUserHandler(r, uc)

	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	port := ":8000"
	log.Println("Listening on port", port)
	srv := &http.Server{
		Addr:              port,
		Handler:           r,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	<-stop
	log.Println("Server shutdown")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
