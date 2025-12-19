# simlab

## Getting started

Run the following command to generate a network topology with each node assigned a bandwidth and a location.
```bash
go run ./cmd/gen-topology/main.go
```

There are also two important things exported from the package.
- The function `simlab.GetTopology` which loads a topology from a file path.
- The global variable `simlab.Lantencies` which contains the latencies across countries.
