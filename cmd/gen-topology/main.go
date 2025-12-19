package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ethp2p/simlab"
)

//go:embed weights.json
var weightsBytes []byte

func main() {
	numNodes := flag.Int("n", 10, "number of nodes")
	degree := flag.Int("d", 4, "degree of the graph (connections per node)")
	superNodeFraction := flag.Float64("super-node-fraction", 0.0, "super-nodes: include 1024Mbit supernodes")
	flag.Parse()

	if *numNodes <= 0 {
		panic(fmt.Sprintf("numNodes <= 0 (%d)", *numNodes))
	}

	var weights map[string]int
	if err := json.Unmarshal(weightsBytes, &weights); err != nil {
		panic("invalid weights")
	}

	t := simlab.GenerateRandomTopology(*numNodes, *degree, weights, *superNodeFraction)

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshalling topology: %v\n", err)
	}
	fmt.Printf("%s", string(data))
}
