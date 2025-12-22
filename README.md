# simlab

## Getting started

Run the following command to generate a network topology with each node assigned a bandwidth and a location.
```bash
go run ./cmd/gen-topology/main.go
```

The JSON format of the generated topology is as follows.
```
type Node struct {
	Country        string `json:"country"`
	UploadBWMbps   int    `json:"uploadBWMbps"`
	DownloadBWMbps int    `json:"downloadBWMbps"`
}

type Topology struct {
	Nodes []Node        `json:"nodes"`
	Edges map[int][]int `json:"edges"`
}
```

There are also two important things exported from the package.
- The function `simlab.GetTopology` which loads a topology from a file path.
- The global variable `simlab.Lantencies` which contains the latencies across countries.

The latencies are hard-coded in `country_latencies.json`.
