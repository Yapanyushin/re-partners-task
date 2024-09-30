package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	web "github.com/Yapanyushin/tabeo-challenge/internal/ui/pack_calculator"
)

func uiCommand(_ *cobra.Command, _ []string) {
	srv := web.NewServer(os.Getenv("API_URL"), os.Getenv("CLIENT_PORT"))

	go func() {
		log.Println("starting service on address", srv.Server.Addr)
		if err := srv.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch
	log.Println("Received shutdown signal. Gracefully stopping...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server stopped")
}
