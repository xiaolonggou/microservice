package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/xiaolonggou/microservice/v1/data"
	"github.com/xiaolonggou/microservice/v1/handlers"
)

func main() {
	l := hclog.Default()
	v := data.NewValidation()
	aph := handlers.NewArtPiece(l, v)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/arts", aph.GetArtPieces)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", aph.UpdateArtPiece)
	putRouter.Use(aph.MiddlewareArtPieceValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/arts", aph.AddArtPiece)
	postRouter.Use(aph.MiddlewareArtPieceValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/art/{id:[0-9]+}", aph.DeleteArtPiece)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	getRouter.Handle("/docs", sh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Info("server starting...")
		l.Debug("serving HTTP...")
		err := s.ListenAndServe()

		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Got signal:", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
