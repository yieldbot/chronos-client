package main

import (
	"fmt"
	"log"

	"github.com/yieldbot/chronos-client/examples/chronos"
)

func main() {
	jobs, err := chronos.Client.Jobs()
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		fmt.Printf("%s\n", j.Name)
	}
}
