package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMapKey(t *testing.T) {
	Convey("Given one map.", t, func() {
		m := make(map[string][]IRouteHandler)
		Convey("When one key doesn't exist", func() {
			c := len(m["a"])
			Convey("its length should be 0", func() {
				So(c, ShouldEqual, 0)
			})
		})
	})

	Convey("Given one map.", t, func() {
		m := make(map[string][]IRouteHandler)
		Convey("When add one key", func() {
			m["a"] = []IRouteHandler{
				NewRoute(API, "/hello", "GET", testHandler),
			}
			Convey("its length should be greater than 0", func() {
				So(len(m["a"]), ShouldEqual, 1)
			})
		})
	})
}

func testHandler(c *gin.Context) {

}
