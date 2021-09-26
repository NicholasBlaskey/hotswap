package main

import (
	"github.com/nicholasblaskey/hotswap/hotswap"
	"os/exec"
)

func main() {
	hotswap.Serve(
		// Directory to server
		"./exampleWasm",
		// Port
		8080,
		// Command to run upon refresh
		exec.Command("cat", "hello", " > out.txt"),
		//exec.Command("go",
		//	"GOOS=js GOARCH=wasm build ./exampleWasm/example.go", "-o ./exampleWasm/o.wasm"),
	)
}
