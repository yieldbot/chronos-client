package main

import (
	"fmt"
	"log"

	"github.com/yieldbot/chronos-client/examples/chronos"
)

func main() {
	var j = `{"schedule": "R/2015-11-09T00:00:00Z/PT24H", "name": "test-1", "epsilon": "PT30M", "command": "echo test1 && sleep 60", "owner": "localhost@localhsot", "async": false}`
	_, err := chronos.Client.AddJob(j)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The job is added\n")
}
