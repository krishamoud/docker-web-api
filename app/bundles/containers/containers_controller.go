package containers

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"github.com/krishamoud/docker-stats/app/common/controller"
	"github.com/krishamoud/docker-stats/app/common/docker"
	"net/http"
)

// ContainersController struct
type ContainersController struct {
	common.Controller
}

// Index func return all containers
func (c *ContainersController) Index(w http.ResponseWriter, r *http.Request) {
	containerList, err := docker.DockerConn.ContainerList(context.Background(), types.ContainerListOptions{})
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		containerList,
		http.StatusOK,
	)
}

// Show a single container
func (c *ContainersController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	containerId := vars["containerId"]
	container, err := docker.DockerConn.ContainerInspect(context.Background(), containerId)
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		container,
		http.StatusOK,
	)
}
