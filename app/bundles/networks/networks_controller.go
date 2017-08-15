package networks

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"github.com/krishamoud/docker-stats/app/common/controller"
	"github.com/krishamoud/docker-stats/app/common/docker"
	"net/http"
)

// NetworksController struct
type NetworksController struct {
	common.Controller
}

// Index func return all networks
func (c *NetworksController) Index(w http.ResponseWriter, r *http.Request) {
	containerList, err := docker.DockerConn.NetworkList(context.Background(), types.NetworkListOptions{})
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

// Show a single network
func (c *NetworksController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	networkId := vars["networkId"]
	network, err := docker.DockerConn.NetworkInspect(context.Background(), networkId, types.NetworkInspectOptions{})
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		network,
		http.StatusOK,
	)
}
