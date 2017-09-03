package volumes

import (
	"context"
	_ "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gorilla/mux"
	"github.com/krishamoud/docker-stats/app/common/controller"
	"github.com/krishamoud/docker-stats/app/common/docker"
	"net/http"
)

// VolumesController struct
type VolumesController struct {
	common.Controller
}

// Index func return all volumes
func (c *VolumesController) Index(w http.ResponseWriter, r *http.Request) {
	volumeList, err := docker.DockerConn.VolumeList(context.Background(), filters.Args{})
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		volumeList,
		http.StatusOK,
	)
}

// Show a single volume
func (c *VolumesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	volumeId := vars["volumeId"]
	volume, err := docker.DockerConn.VolumeInspect(context.Background(), volumeId)
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		volume,
		http.StatusOK,
	)
}
