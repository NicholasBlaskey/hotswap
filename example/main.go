package main

import (
	"github.com/nicholasblaskey/hotswap/hotswap"
)

func main() {
	hotswap.Serve(
		// Directory to server
		"./exampleWasm",
		// Port
		8080,
		// Command to run upon refresh
		"GOOS=js GOARCH=wasm go build exampleWasm/example.go -o exampleWasm/o.wasm",
	)
}
