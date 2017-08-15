package containers_test

import (
	"github.com/krishamoud/docker-stats/app/bundles/containers"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllerSpec(t *testing.T) {
	Convey("Given a running server and controller instance", t, func() {
		c := containers.ContainersController{}
		mux := http.NewServeMux()
		server := httptest.NewServer(mux)
		defer server.Close()
		Convey("When SendJSON is called from handler with 200", func() {
			mux.HandleFunc("/", c.Index)
			resp, err := http.Get(server.URL + "/test1")
			Convey("Then response should be 200 with correct JSON", func() {
				if err != nil {
					t.Fatal(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})
}
