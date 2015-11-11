package chronos

import (
	"os"

	"github.com/yieldbot/chronos-client"
)

var (
	// URL is the Chronos url
	URL string

	// Client is the Chronos client
	Client client.Client
)

func init() {
	// Init the Chronos client
	URL = "localhost:8080"
	if os.Getenv("CHRONOS_URL") != "" {
		URL = os.Getenv("CHRONOS_URL")
	}
	Client = client.Client{URL: URL}
}
