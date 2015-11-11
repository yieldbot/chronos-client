package main

import (
	"log"

	"github.com/yieldbot/chronos-client/examples/chronos"
)

func main() {
	if err := chronos.Client.PrintJobs(false); err != nil {
		log.Fatal(err)
	}
}
