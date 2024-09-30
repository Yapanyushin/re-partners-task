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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Yapanyushin/tabeo-challenge/api/proto"
	web "github.com/Yapanyushin/tabeo-challenge/internal/ui/pack_calculator"
)

func uiCommand(_ *cobra.Command, _ []string) {

	// Set up a connection to the server
	conn, err := grpc.NewClient(os.Getenv("API_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cant establish connection %s", err.Error())
	}

	defer conn.Close()

	srv := web.NewServer(proto.NewPackCalculatorClient(conn), os.Getenv("CLIENT_PORT"))

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
