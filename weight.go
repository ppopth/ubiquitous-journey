package simlab

import (
	_ "embed"
	"encoding/json"
	"log"
)

var Weights map[string]int

//go:embed weights.json
var weightsBytes []byte

func init() {
	Weights = make(map[string]int)

	if err := json.Unmarshal(weightsBytes, &Weights); err != nil {
		log.Fatalf("Failed to parse weight data: %v", err)
	}
}
