package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		osName := runtime.GOOS

		// Try running curl --version
		cmd := exec.Command("aws", "--version")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Fprintf(w, "Running on OS: %s\n\n", osName)
			fmt.Fprintf(w, "aws check: ❌ not available or failed to run\n")
			fmt.Fprintf(w, "error: %v\n", err)
		} else {
			fmt.Fprintf(w, "Running on OS: %s\n\n", osName)
			fmt.Fprintf(w, "aws check: ✅ available and working!\n")
			fmt.Fprintf(w, "Output:\n%s\n", output)
		}
	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
