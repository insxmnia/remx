package server

import (
	"os"
	"os/signal"
	"remx/internal/ui"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	router := echo.New()
	go func() {
		<-sigs
		router.Close()
		done <- true
	}()
	router.HideBanner = true
	router.HidePort = true
	router.Use(middleware.Recover())
	router.Use(CLILogger)
	router.Any("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"received": true,
		})
	})
	ui.Colours["secondary"].Faith.Printf("Listening for requests on %s:%s...\n", ui.Colours["primary"].Faith.Sprint("0.0.0.0"), ui.Colours["primary"].Faith.Sprint(port))
	router.Start(":" + port)
}
