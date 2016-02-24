package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"encoding/xml"
	"path/filepath"
	_ "github.com/deluan/gosonic/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"github.com/deluan/gosonic/controllers/responses"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.." + string(filepath.Separator))))
	beego.TestBeegoInit(appPath)
}

// TestGet is a sample to run an endpoint test
func TestPing(t *testing.T) {
	r, _ := http.NewRequest("GET", "/rest/ping.view", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPing", fmt.Sprintf("Code[%d]\n%s", w.Code, w.Body.String()))

	Convey("Subject: Ping Endpoint\n", t, func() {
		Convey("Status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The result should not be empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("The result should be a valid ping response", func() {
			v := responses.Subsonic{}
			xml.Unmarshal(w.Body.Bytes(), &v)
			So(v.Status, ShouldEqual, "ok")
			So(v.Version, ShouldEqual, "1.0.0")
		})

	})
}

