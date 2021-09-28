package hotswap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func Serve(directory string, port int, cmd *exec.Cmd) {
	log.Printf("Starting server the directory %s port=%d\n\n", directory, port)

	http.HandleFunc("/compile", func(w http.ResponseWriter, _ *http.Request) {
		// Copy command
		cmdCopy := *cmd

		log.Printf("Executing the command: %s\n", cmdCopy.String())
		out, err := cmdCopy.CombinedOutput()

		log.Printf("Output of command: %s\n", string(out))
		if err != nil {
			log.Print(err)
		}
		fmt.Println()

		io.WriteString(w, string(out))
	})

	http.Handle("/",
		http.FileServer(http.Dir(directory)),
	)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
