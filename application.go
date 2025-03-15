package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	serverRandom := rand.IntN(1000)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World " + strconv.Itoa(serverRandom)))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":5000", nil)
}
