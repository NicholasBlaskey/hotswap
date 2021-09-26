package hotswap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func Serve(directory string, port int, cmd *exec.Cmd) {
	log.Printf("Starting server the directory %s port=%d\n", directory, port)

	http.HandleFunc("/compile", func(w http.ResponseWriter, _ *http.Request) {
		log.Printf("Executing the command: %s\n", cmd)
		err := cmd.Run()
		if err != nil {
			log.Print(err)
		}

		io.WriteString(w, "")
	})

	http.Handle("/",
		http.FileServer(http.Dir(directory)),
	)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

/*
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
*/

/*
	hotswap.Serve(
		// Directory to server
		"./exampleWasm",
		// Port
		8080,
		// Command to run upon refresh
		"GOOS=js GOARCH=wasm go build exampleWasm/example.go -o exampleWasm/o.wasm",
	)
*/
