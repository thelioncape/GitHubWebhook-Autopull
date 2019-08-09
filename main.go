package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// X52nqgk3wZG9A3oXa55mqLM71jJQPeRhMuVDC8Ua

func main() {
	wh := func(w http.ResponseWriter, _ *http.Request) {

		f, err := os.Open("location.txt")
		if err != nil {
			log.Fatal(err)
		}

		s := bufio.NewScanner(f)
		s.Scan()
		location := s.Text()

		cmd := exec.Command("git", "pull")
		cmd.Dir = location

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/X52nqgk3wZG9A3oXa55mqLM71jJQPeRhMuVDC8Ua", wh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
