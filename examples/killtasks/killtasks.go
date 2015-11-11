package main

import (
	"fmt"
	"log"

	"github.com/yieldbot/chronos-client/examples/chronos"
)

func main() {
	_, err := chronos.Client.KillJobTasks("test-1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("test-1 job tasks are killed\n")
}
