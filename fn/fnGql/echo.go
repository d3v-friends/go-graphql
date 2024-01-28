package fnGql

import (
	"context"
	"errors"
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

const CtxHeader = "CTX_HEADER"

var ErrNotFoundHeader = errors.New("not found header")

func NewEcho() (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
	return
}

func Launch(e *echo.Echo, port string) {
	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	fnPanic.On(e.Start(port))
}

func HeaderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var ctx = c.Request().Context()
			ctx = context.WithValue(ctx, CtxHeader, c.Request().Header)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}

func GetHeader(ctx context.Context) (header http.Header, err error) {
	var isOk bool
	if header, isOk = ctx.Value(CtxHeader).(http.Header); !isOk {
		err = ErrNotFoundHeader
		return
	}
	return
}

func GetHeaderP(ctx context.Context) (header http.Header) {
	var err error
	if header, err = GetHeader(ctx); err != nil {
		panic(err)
	}
	return
}
