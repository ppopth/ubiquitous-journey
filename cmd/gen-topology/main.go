package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/ethp2p/simlab"
)

func main() {
	numNodes := flag.Int("n", 10, "number of nodes")
	degree := flag.Int("d", 4, "degree of the graph (connections per node)")
	superNodeFraction := flag.Float64("super-node-fraction", 0.0, "super-nodes: include 1024Mbit supernodes")
	flag.Parse()

	if *numNodes <= 0 {
		panic(fmt.Sprintf("numNodes <= 0 (%d)", *numNodes))
	}

	t := simlab.GenerateRandomTopology(*numNodes, *degree, *superNodeFraction)

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshalling topology: %v\n", err)
	}
	fmt.Printf("%s", string(data))
}
