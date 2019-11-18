package handler

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"wps.ktkt.com/app2017/elastic_query-test/pkg/response"
)


type ArgsInfo struct {

}

// meta数据，换取用户token
// meta 系统【安卓，ios， pc， iPad，h5 】，系统版本【】， 软件版本【】， 用户ID，
func Token(c echo.Context) error {
	glog.Info("[Token] called.")

	args := ArgsInfo{}

	err := c.Bind(&args)
	if err != nil {
		return ErrorParams
	}

	glog.Infof("[Token] args : %#v \n", args)

	// Generate encoded token and send it as response.

	data := "ok"

	return c.JSON(http.StatusOK, response.OKResponse(data))
}
