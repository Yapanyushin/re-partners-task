package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/Yapanyushin/tabeo-challenge/api/proto"
	"github.com/Yapanyushin/tabeo-challenge/internal/app"
	"github.com/Yapanyushin/tabeo-challenge/internal/server/pack_calculator"
)

func serveCommand(_ *cobra.Command, _ []string) {
	config := readConfig()

	ln, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalf("comand open port error %s", err)
	}

	s := grpc.NewServer()

	pc, err := app.NewPackCalculator(config.PackSizes)
	if err != nil {
		log.Fatalf("app.NewPackCalculator error: %s", err.Error())
	}

	proto.RegisterPackCalculatorServer(s, pack_calculator.NewServer(pc))

	healthpb.RegisterHealthServer(s, health.NewServer())

	go func() {
		log.Println("starting service")
		if err := s.Serve(ln); err != nil {
			log.Fatalf("error serve %s", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	<-ch
	log.Println("Received shutdown signal. Gracefully stopping...")

	s.GracefulStop()

	log.Println("Server stopped")
}
