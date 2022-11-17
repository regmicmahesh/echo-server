package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		map_ := r.Header.Clone()

		map_["Uri"] = []string{r.RequestURI}

		map_["Method"] = []string{r.Method}

		b, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		map_["Body"] = []string{string(b)}

		map_["Env"] = os.Environ()

		headersJSON, _ := json.Marshal(map_)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(headersJSON)

	})

	if port, found := os.LookupEnv("PORT"); !found {
		startServer("8000", mux)
	} else {
		startServer(port, mux)
	}

}

func startServer(port string, mux *http.ServeMux) {
	println("Listening on port " + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}
