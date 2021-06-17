package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wiltonassessment/internal/cond"
	"wiltonassessment/internal/config"
)

func Run(path string) error {
	cfg, err := config.GetConfig(path)
	if err != nil {
		log.Fatalf("Error occured during reading configuration file: %v\n", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", cond.GetInputData)
	srv := NewServer(cfg, mux)

	go func() {
		log.Println("Server started", srv.httpServer.Addr)
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error occured during starting the server: %v\n", err)
		}
	}()
	<-stop

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("Error occured during stopping the server: %v\n", err)
	}
	return nil
}
