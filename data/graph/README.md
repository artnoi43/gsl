# Package `graph`

This package provides basic building blocks for working with graphs.

There's the interface `GenericGraph`, which all other graph implementations must implement.

The algorithm code then operates on `GenericGraph`. You can easily implement `GenericGraph`
with your own types, as shown in [this Dijkstra example](./assets/example_flight_path/main.go).
