package main

import (
	"log"
	"net/http"
	"sync"

	v1 "arctair.com/quarky/v1"
)

var (
	sha1    string
	version string
)

// StartHTTPServer ...
func StartHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{
		Addr: ":5000",
		Handler: v1.NewRouter(
			v1.NewRolloutController(v1.NewDeployments(), &v1.LoggerConsole{}),
			v1.NewVersionController(
				v1.NewBuild(sha1, version),
			),
		),
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
