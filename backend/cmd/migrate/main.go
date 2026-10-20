package main

import (
	"context"
	"fmt"
	"os"

	"catalina-support/backend/shared/migrations"
)

func main() {
	if err := migrations.Run(context.Background(), migrations.ConfigFromEnv()); err != nil {
		fmt.Fprintf(os.Stderr, "migration failed: %v\n", err)
		os.Exit(1)
	}
}
