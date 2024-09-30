package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Port      string
	PackSizes []int32
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "bookings",
		Short: "service serves concurrent booking for shared launchpads",
	}

	cmdServe := &cobra.Command{
		Use:   "serve",
		Short: "Starts grpc service",
		Run:   serveCommand,
	}

	rootCmd.AddCommand(
		cmdServe,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Errorf("error start execute: %w", err))
	}

}

func readConfig() Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s", err.Error())
	}

	s := viper.GetIntSlice("pack_sizes")
	sizes := make([]int32, len(s))

	for i, size := range s {
		sizes[i] = int32(size)
	}

	return Config{
		Port:      os.Getenv("SERVER_PORT"),
		PackSizes: sizes,
	}
}
