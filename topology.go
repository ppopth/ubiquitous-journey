package simlab

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"slices"
)

type Node struct {
	Country        string `json:"country"`
	UploadBWMbps   int    `json:"uploadBWMbps"`
	DownloadBWMbps int    `json:"downloadBWMbps"`
}

// Topology represents the complete network topology with nodes and edges
type Topology struct {
	Nodes []Node        `json:"nodes"`
	Edges map[int][]int `json:"edges"`
}

func GetTopology(fileName string) *Topology {
	f, err := os.Open(fileName)
	if err != nil {
		panic("failed to read scenario " + fileName)
	}
	d := json.NewDecoder(f)
	t := &Topology{}
	if err := d.Decode(t); err != nil {
		panic("failed to read scenario: " + fileName + "err: " + err.Error())
	}
	return t
}

type countrySelector struct {
	countries         []string
	cumulativeWeights []int
	totalWeight       int
}

func newCountrySelector(weights map[string]int) countrySelector {
	countries := make([]string, 0, len(weights))
	cumulativeWeights := make([]int, 0, len(weights))

	total := 0
	for country, weight := range weights {
		if weight <= 0 {
			continue
		}
		countries = append(countries, country)
		total += weight
		cumulativeWeights = append(cumulativeWeights, total)
	}
	return countrySelector{
		countries:         countries,
		cumulativeWeights: cumulativeWeights,
		totalWeight:       total,
	}
}

// randCountry randomly selects a country based on pre-computed weights
func (cs *countrySelector) randCountry() string {
	r := rand.IntN(cs.totalWeight)
	index, _ := slices.BinarySearch(cs.cumulativeWeights, r)
	return cs.countries[index]
}

// GenerateRandomTopology creates a topology with the specified number of nodes and degree
func GenerateRandomTopology(numNodes, degree int, superNodeFraction float64) Topology {
	t := Topology{
		Nodes: make([]Node, numNodes),
		Edges: make(map[int][]int),
	}

	selector := newCountrySelector(Weights)

	for i := range numNodes {
		uploadBW, downloadBW := getNodeBandwidth(i, superNodeFraction)
		t.Nodes[i] = Node{
			Country:        selector.randCountry(),
			UploadBWMbps:   uploadBW,
			DownloadBWMbps: downloadBW,
		}
		t.Edges[i] = []int{}
	}

	// first generate a directed graph
	for i := 1; i < numNodes; i++ {
		v := rand.IntN(i) // connect to one of the prior already connected nodes
		addEdge(t.Edges, i, v)
	}

	maxAttempts := numNodes * 10
	// add edges for reaching the desired degree
	for u := range numNodes {
		attempts := 0
		for len(t.Edges[u]) < degree && attempts < maxAttempts {
			v := rand.IntN(numNodes)
			if v == u || slices.Contains(t.Edges[u], v) || len(t.Edges[v]) >= degree {
				attempts++
				continue
			}
			addEdge(t.Edges, u, v)
			attempts++
		}
	}
	return t
}

const (
	uploadBWSuperNode   = 1024
	downloadBWSuperNode = 1024

	uploadBWBlockBuilder   = 50
	downloadBWBlockBuilder = 100

	uploadBWFullNode   = 15
	downloadBWFullNode = 50

	uploadBWAttester   = 25
	downloadBWAttester = 50

	uploadBWFullNodeOld   = 50
	downloadBWFullNodeOld = 50
)

// getNodeBandwidth determines bandwidth based on node type
func getNodeBandwidth(nodeId int, superNodeFraction float64) (uploadBW, downloadBW int) {
	if superNodeFraction > 0.0001 && nodeId == 0 {
		return uploadBWSuperNode, downloadBWSuperNode
	}
	if nodeId == 0 {
		return uploadBWBlockBuilder, downloadBWBlockBuilder
	}
	x := rand.IntN(1000)
	if x < int(1000*superNodeFraction) {
		return uploadBWSuperNode, downloadBWSuperNode
	}
	return uploadBWFullNodeOld, downloadBWFullNodeOld
}

func addEdge(edges map[int][]int, a, b int) bool {
	if !slices.Contains(edges[a], b) && !slices.Contains(edges[b], a) {
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
		return true
	}
	return false
}
