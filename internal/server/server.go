package server

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/golang/glog"
	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"wps.ktkt.com/app2017/elastic_query-test/configs"
	"wps.ktkt.com/app2017/elastic_query-test/internal/server/handler"
	"wps.ktkt.com/app2017/elastic_query-test/pkg/response"

	//"wps.ktkt.com/app2017/elastic_query-test/configs"

	lerr "wps.ktkt.com/app2017/elastic_query-test/pkg/error"
	"wps.ktkt.com/app2017/elastic_query-test/pkg/rotate"
)

func StartHttpServer(conf *configs.WebConfig, ver string) {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	if len(conf.AccessLogPath) > 0 {
		dir := path.Dir(conf.AccessLogPath)

		err := os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			glog.Error("mkdir failed.", dir, err)
		}

		fileLog := rotate.New(conf.AccessLogPath)


		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: io.MultiWriter(os.Stdout, fileLog)}))
	} else {
		e.Use(middleware.Logger())
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, ver)
	})

	// add router
	handler.Router(e)

	glog.Info("start http server: ", conf.Listen)

	e.Logger.Fatal(e.Start(conf.Listen))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	var msg interface{}

	msg = "服务器错误"

	switch err.(type) {
	case *echo.HTTPError:
		he := err.(*echo.HTTPError)
		code = he.Code
		msg = he.Message

	case lerr.Err, *lerr.Err:
		he := err.(*lerr.Err)
		code = he.Code
		msg = he.Message

	default:
		glog.Errorf("unknown ERR %T %v", err, err)
	}

	// error response

	c.JSON(http.StatusOK, response.ErrorResponse(code, msg))

	//if c.Request().Header().Get("X-Requested-With") == "XMLHttpRequest" {
	//	// ajax 请求 返回json
	//	c.JSON(http.StatusOK, api.RetErr(code, msg))
	//	return
	//}

	return
}
