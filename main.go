package main

import (
	"fmt"

	// Docker client
	"github.com/docker/docker/client"

	// Router
	"github.com/krishamoud/docker-stats/app/common/router"

	// Common code
	"github.com/krishamoud/docker-stats/app/common/docker"

	"log"
	"net/http"
)

func main() {
	var err error
	docker.DockerConn, err = client.NewEnvClient()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// Handle all requests with gorilla/mux
	http.Handle("/", router.Router())

	// Listen on port 9090
	log.Println("Server listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
