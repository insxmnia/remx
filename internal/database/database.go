package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"remx/pkg/slogger"
	"runtime"
)

func (c *DatabaseCredentials) Combine(key string) string {
	return fmt.Sprintf("%s%s%s%s", c.EndpointType, c.Endpoint, key, c.Token)
}

func (d *Database) Initialize() error {
	if d.db != nil {
		slogger.Error("database", "initialization failed due to an instance already available")
		return fmt.Errorf("initialization failed due to an instance already available")
	}

	db, err := sql.Open(d.Options.Driver, d.Credentials.Combine("?authToken="))
	if err != nil {
		slogger.Fatal("database", "failed to open sql connection", "error", err)
	}
	if d.Options == nil {
		d.Options = &Defaults
	}

	db.SetMaxOpenConns(d.Options.MaxOpenConns)
	db.SetConnMaxIdleTime(d.Options.MaxIdleTime)

	if err := db.PingContext(context.Background()); err != nil {
		slogger.Fatal("database", "ping failed to endpoint", "error", err)
	}

	d.db = db
	slogger.Info("database", "successfully initialized connection and instances")
	return nil
}

func (d *Database) guard(callback func()) {
	pointer := reflect.ValueOf(callback).Pointer()
	name := runtime.FuncForPC(pointer).Name()
	if d.db == nil {
		slogger.Error("database", "guard failed due to uninitialized database instance", "callback", name)
		return
	}
	if err := d.db.Ping(); err != nil {
		slogger.Error("database", "guard failed due to database ping error", "callback", name, "error", err)
		return
	}
	callback()
}
