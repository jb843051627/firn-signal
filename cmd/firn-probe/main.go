package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jb843051627/firn-signal/internal/api"
	"github.com/jb843051627/firn-signal/internal/service"
	"github.com/jb843051627/firn-signal/internal/store"
)

func main() {
	path := os.Getenv("FIRN_PROBE_DB")
	if path == "" {
		log.Fatal("FIRN_PROBE_DB is required")
	}
	repository, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close()
	lab := service.NewLab(repository)
	defer lab.Close()
	server := httptest.NewServer(api.New(lab))
	defer server.Close()
	response, err := http.Get(server.URL + "/healthz")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("health status: %s", response.Status)
	}
	var health struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(response.Body).Decode(&health); err != nil {
		log.Fatal(err)
	}
	if health.Path != path {
		log.Fatalf("database path %q != %q", health.Path, path)
	}
	if _, err := os.Stat(path); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("firn probe passed: %s\n", path)
}
