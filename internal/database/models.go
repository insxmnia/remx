package database

import (
	"database/sql"
	"time"
)

type DatabaseCredentials struct {
	EndpointType string
	Endpoint     string
	Token        string
}
type DatabaseOptions struct {
	MaxPagingSize int
	MaxOpenConns  int
	MaxIdleTime   time.Duration
	Driver        string
}
type Database struct {
	db          *sql.DB
	Credentials DatabaseCredentials
	Options     *DatabaseOptions
}
