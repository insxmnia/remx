package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"remx/internal/ui"
	"time"

	"github.com/labstack/echo/v4"
)

func CLILogger(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		err := handler(ctx)
		end := time.Now()

		latency := end.Sub(start)

		method := ctx.Request().Method
		path := ctx.Request().URL.Path
		status := ctx.Response().Status
		ip := ctx.RealIP()
		body := ctx.Request().Body
		read, err := io.ReadAll(body)
		if err != nil {
			read = []byte("")
		}

		var headers []map[string][]string
		for key, val := range ctx.Request().Header {
			headers = append(headers, map[string][]string{key: val})
		}
		header, err := json.Marshal(headers)
		if err != nil {
			header = []byte("")
		}

		ui.Colours["secondary"].Faith.Println(fmt.Sprintf("%s %s |%s| %s | %s %s | query: [%s] | body: [%s] | headers: [%s] ",
			end.Format("2006/01/02 15:04:05"),
			ip,
			ui.Colours["primary"].Faith.Sprint(status)+ui.Colours["secondary"].ANSII,
			latency,
			method,
			path,
			ctx.Request().URL.RawQuery,
			base64.StdEncoding.EncodeToString(read),
			string(header),
		))

		return err
	}
}
