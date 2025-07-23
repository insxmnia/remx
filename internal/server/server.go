package server

import (
	"remx/internal/ui"

	"github.com/eiannone/keyboard"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(port string) {
	keyboard.Open()
	defer keyboard.Close()
	router := echo.New()
	router.HideBanner = true
	router.HidePort = true
	router.Use(middleware.Recover())
	router.Use(CLILogger)
	router.Any("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"received": true,
		})
	})
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			if char == 'q' || key == keyboard.KeyEsc {
				router.Close()
				break
			}
		}
	}()
	ui.Colours["secondary"].Faith.Println("Use 'q' or 'esc' to quit...")
	ui.Colours["secondary"].Faith.Printf("Listening for requests on %s:%s...\n", ui.Colours["primary"].Faith.Sprint("0.0.0.0"), ui.Colours["primary"].Faith.Sprint(port))

	router.Start(":" + port)
}
