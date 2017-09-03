package images

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gorilla/mux"
	"github.com/krishamoud/docker-stats/app/common/controller"
	"github.com/krishamoud/docker-stats/app/common/docker"
	"net/http"
)

// ImagesController struct
type ImagesController struct {
	common.Controller
}

// Index func return all images
func (c *ImagesController) Index(w http.ResponseWriter, r *http.Request) {
	imageList, err := docker.DockerConn.ImageList(context.Background(), types.ImageListOptions{})
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		imageList,
		http.StatusOK,
	)
}

// Show a single image
func (c *ImagesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["imageId"]
	image, _, err := docker.DockerConn.ImageInspectWithRaw(context.Background(), imageId)
	if c.CheckError(err, http.StatusInternalServerError, w) {
		return
	}
	c.SendJSON(
		w,
		r,
		image,
		http.StatusOK,
	)
}
