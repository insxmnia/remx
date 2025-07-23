package database

import "time"

var Defaults DatabaseOptions = DatabaseOptions{
	MaxPagingSize: 15000,
	MaxOpenConns:  10,
	MaxIdleTime:   time.Second * 2,
	Driver:        "turso-remote",
}
