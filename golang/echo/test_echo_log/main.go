package main

import (
    "fmt"
    "net/http"
	"os"
	"strings"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type Response struct {
    Status  int
    Message string
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("%v\n", c.Cookies())
    fmt.Printf("Request Body: %v\n", string(reqBody))
    fmt.Printf("Response Body: %v\n", string(resBody))
}


func logFormat() string {
	// Refer to https://github.com/tkuchiki/alp
	formatArray := []string{
		"time:${time_rfc3339}\t",
		"host:${remote_ip}\t",
		"forwardedfor:${header:x-forwarded-for}\t",
		"req:-\t",
		"status:${status}\t",
		"method:${method}\t",
		"uri:${uri}\t",
		"size:${bytes_out}\t",
		"referer:${referer}\t",
		"ua:${user_agent}\t",
		"reqtime_ns:${latency}\t",
		"cache:-\t",
		"runtime:-\t",
		"apptime:-\t",
		"vhost:${host}\t",
		"reqtime_human:${latency_human}\t",
		"x-request-id:${id}\t",
		"host:${host}\t",
		"header:${header}\t",
		"query:${query}\t",
		"form:${form}\t",
		"cookie:${cookir}\t",
	}
	format := strings.Join(formatArray, "")
	return format
}

func main() {
    e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	}))
    e.Use(middleware.BodyDump(bodyDumpHandler))

    e.HTTPErrorHandler = func(err error, c echo.Context) {
        if c.Response().Committed {
            return
        }
        if he, ok := err.(*echo.HTTPError); ok {
            c.JSON(he.Code, Response{
                Status:  he.Code,
                Message: he.Error(),
            })
        }
    }

    e.POST("/yorimoi", func(c echo.Context) error {
        var crew []string
        c.Bind(&crew)
        if len(crew) < 4 {
            return echo.NewHTTPError(http.StatusBadRequest, "Going by 4 people is top priority!")
        }

        return c.JSON(http.StatusOK, Response{
            Status:  http.StatusOK,
            Message: "They got to the Antarctica!",
        })
    })

    e.Logger.Fatal(e.Start(":1323"))

}
