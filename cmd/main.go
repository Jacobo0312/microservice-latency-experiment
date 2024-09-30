package main

import (
	"fmt"

	"github.com/Jacobo0312/microservice-latency-experiment/cmd/app/container"
	"github.com/Jacobo0312/microservice-latency-experiment/internal/rest"
)

func main() {

	readerHandler := container.Reader()

	if err := rest.API(readerHandler).Run(); err != nil {
		panic(fmt.Errorf("error initializing api: %w", err))
	}

}
