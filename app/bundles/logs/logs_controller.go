package logs

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/krishamoud/docker-stats/app/common/controller"
	"github.com/krishamoud/docker-stats/app/common/docker"
	"net/http"
)

// LogsController struct
type LogsController struct {
	common.Controller
}

// Websocket upgrade to push logs
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Stream the logs to the client
func (c *LogsController) Show(w http.ResponseWriter, r *http.Request) {
	// get the variables from mux
	vars := mux.Vars(r)
	containerId := vars["containerId"]

	// open the read stream from docker
	reader, err := docker.DockerConn.ContainerLogs(context.Background(), containerId,
		types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Timestamps: true,
			Details:    false,
		})

	// if the reader fails return an internal server error
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}

	// close the reader after we're done with it
	defer reader.Close()

	// upgrade the connection for websockets
	conn, err := upgrader.Upgrade(w, r, nil)
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}

	// create a new scanner and read from it until it closes
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := []rune(scanner.Text())
		err = conn.WriteMessage(websocket.TextMessage, []byte(string(text)))
		if c.CheckError(err, http.StatusInternalServerError, w) {
			return
		}
	}
	if err = scanner.Err(); err != nil {
		if c.CheckError(scanner.Err(), http.StatusInternalServerError, w) {
			return
		}
	}
}
