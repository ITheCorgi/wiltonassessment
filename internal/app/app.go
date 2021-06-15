package app

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"wiltonassessment/internal/cond"
	"wiltonassessment/internal/config"
)

func Run(path string) error {
	errs := make(chan error)

	cfg, err := config.GetConfig(path)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", cond.GetInputData)
	srv := NewServer(cfg, mux)

	terminating := make(chan os.Signal, 1)
	signal.Notify(terminating, syscall.SIGTERM, syscall.SIGINT)
	<-terminating

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
	}()

	for err := range errs {
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
