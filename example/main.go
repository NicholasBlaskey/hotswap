package main

import (
	"github.com/nicholasblaskey/hotswap/hotswap"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "build", "-o=./exampleWasm/o.wasm", "./exampleWasm/example.go") //, "-o ./exampleWasm/o.wasm")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")

	hotswap.Serve(
		// Directory to server
		"./exampleWasm",
		// Port
		8080,
		// Command to run upon refresh
		//exec.Command("echo", "hello", " > out.txt"),
		cmd,
	)
}
