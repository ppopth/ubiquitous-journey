package simlab

import (
	_ "embed"
	"encoding/json"
	"log"
	"time"
)

var Latencies map[string]map[string]int

const DefaultLatency = 100 * time.Millisecond // highly unscientific value

//go:embed country_latencies.json
var countryLatenciesBytes []byte

func init() {
	Latencies = make(map[string]map[string]int)

	if err := json.Unmarshal(countryLatenciesBytes, &Latencies); err != nil {
		log.Fatalf("Failed to parse latency data: %v", err)
	}
}
