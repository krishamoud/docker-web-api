package router

import (
	// mux router
	"github.com/gorilla/mux"

	// Middleware chaining
	"github.com/justinas/alice"

	// Resources
	"github.com/krishamoud/docker-stats/app/bundles/containers"
	"github.com/krishamoud/docker-stats/app/bundles/images"
	"github.com/krishamoud/docker-stats/app/bundles/logs"
	"github.com/krishamoud/docker-stats/app/bundles/networks"
	"github.com/krishamoud/docker-stats/app/bundles/volumes"

	// common middleware
	"github.com/krishamoud/docker-stats/app/common/middleware"

	"net/http"
)

func Router() *mux.Router {
	// Mux Router declaration
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	// Controllers declaration
	cc := &containers.ContainersController{}
	lc := &logs.LogsController{}
	nc := &networks.NetworksController{}
	ic := &images.ImagesController{}
	vc := &volumes.VolumesController{}

	// middleware chaining
	commonHandlers := alice.New(middleware.LoggingHandler, middleware.RecoverHandler)

	// Container Information Routes
	s.Handle("/containers", commonHandlers.ThenFunc(cc.Index)).Methods("GET")
	s.Handle("/containers/{containerId}", commonHandlers.ThenFunc(cc.Show)).Methods("GET")

	// Network Information Routes
	s.Handle("/networks", commonHandlers.ThenFunc(nc.Index)).Methods("GET")
	s.Handle("/networks/{networkId}", commonHandlers.ThenFunc(nc.Show)).Methods("GET")

	// Image Information Routes
	s.Handle("/images", commonHandlers.ThenFunc(ic.Index)).Methods("GET")
	s.Handle("/images/{imageId}", commonHandlers.ThenFunc(ic.Show)).Methods("GET")

	// Volume Information Routes
	s.Handle("/volumes", commonHandlers.ThenFunc(vc.Index)).Methods("GET")
	s.Handle("/volumes/{volumeId}", commonHandlers.ThenFunc(vc.Show)).Methods("GET")

	// Container Logs Route
	// Because we send the logs over websockets we can't run normal middleware
	s.HandleFunc("/logs/{containerId}", lc.Show).Methods("GET")

	// Naked route: only being used for testing purposes at the moment
	// change home.html to get logs for a certain container
	r.Handle("/", commonHandlers.ThenFunc(serveHome)).Methods("GET")

	return r
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}
