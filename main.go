package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	sha1    string
	version string
)

// Build ...
type Build interface {
	getSha1() string
	getVersion() string
}

// BuildVars ...
type BuildVars struct{}

func (b *BuildVars) getSha1() string {
	return sha1
}

func (b *BuildVars) getVersion() string {
	return version
}

// StartHTTPServer ...
func StartHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{
		Addr: ":5000",
		Handler: &Router{
			&BuildVars{},
			NewDeployments(),
			&LoggerConsole{},
		},
	}

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return server
}

func main() {
	serverExit := &sync.WaitGroup{}
	serverExit.Add(1)
	StartHTTPServer(serverExit)
	serverExit.Wait()
}
