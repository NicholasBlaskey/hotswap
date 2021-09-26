package hotswap

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Serve(directory string, port int, cmd string) {
	log.Printf("Starting server the directory %s port=%d\n", directory, port)

	http.HandleFunc("/compile", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
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
