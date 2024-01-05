package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/safciplak/capila/src/apm"
	"github.com/safciplak/capila/src/cli"
)

// main
func main() {
	ctx := context.Background()
	err := apm.TraceError(ctx, godotenv.Load())

	if err != nil {
		log.Print("Failed to load the .env")
	}

	cli.Run(os.Args)
}
