package main

import (
	"os"
	"remx/internal/app"
	"remx/internal/config"
	"remx/internal/database"
	"remx/internal/inmemory"
	"remx/pkg/slogger"
	"time"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		slogger.Fatal("entry", "failed to fetch current working directory", "error", err)
	}
	vpr, err := config.ReadConfig(config.ConfigParams{
		Path: []string{"./config", "/etc/conf", cwd + "/config", "./"},
		Name: "config",
	})
	if err != nil {
		slogger.Fatal("entry", "failed to read configuration", "error", err)
	}
	mapped := config.MapConfig(vpr)
	inmemory.CF = mapped

	database := &database.Database{
		Credentials: database.DatabaseCredentials{
			EndpointType: mapped.Database.EndpointType,
			Endpoint:     mapped.Database.Endpoint,
			Token:        mapped.Database.Token,
		},
		Options: &database.DatabaseOptions{
			Driver:        mapped.Database.Driver,
			MaxPagingSize: 15000,
			MaxOpenConns:  10,
			MaxIdleTime:   time.Second * 2,
		},
	}
	if err := database.Initialize(); err != nil {
		slogger.Fatal("entry", "database error", "error", err)
	}
	inmemory.DB = database
	app.Entry()
}
