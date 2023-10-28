package main

import (
	"log"

	"ddd-sample/interface/signals"
)

func main() {
	ctx := signals.SetupSignalHandler()
	if err := Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
