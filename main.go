package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	var key string

	wh := func(w http.ResponseWriter, _ *http.Request) {

		f, err := os.Open("location.txt")
		if err != nil {
			log.Fatal(err)
		}

		s := bufio.NewScanner(f)
		s.Scan()
		location := s.Text()
		s.Scan()
		key = s.Text()

		cmd := exec.Command("git", "pull")
		cmd.Dir = location
		cmd.Run()
	}

	http.HandleFunc("/"+key, wh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
