package tests

import (
	_ "hongId/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"hongID/controllers"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestTelAuthCode1(t *testing.T) {
	r, _ := http.NewRequest("POST", "/v1/hongId/telAuthCode", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.BeeLogger.Trace("testing TestTelAuthCode1 Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 400", func() {
			So(w.Code, ShouldEqual, 400)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestTelAuthCode2(t *testing.T) {

	memberTel := &controllers.MemberTel{
		Tel: "132959",
	}
	reqBytes, _ := json.Marshal(memberTel)

	r, _ := http.NewRequest("POST", "/v1/hongId/telAuthCode", bytes.NewReader(reqBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.BeeLogger.Trace("testing TestTelAuthCode2 Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 400", func() {
			So(w.Code, ShouldEqual, 400)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestTelAuthCode3(t *testing.T) {

	memberTel := &controllers.MemberTel{
		Tel: "18610889275",
	}
	reqBytes, _ := json.Marshal(memberTel)

	r, _ := http.NewRequest("POST", "/v1/hongId/telAuthCode", bytes.NewReader(reqBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.BeeLogger.Trace("testing TestTelAuthCode3 Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestTelAuthCode4(t *testing.T) {

	memberTel := &controllers.MemberTel{
		Tel: "15281075582",
	}
	reqBytes, _ := json.Marshal(memberTel)

	r, _ := http.NewRequest("POST", "/v1/hongId/telAuthCode", bytes.NewReader(reqBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.BeeLogger.Trace("testing TestTelAuthCode4 Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
