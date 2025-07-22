package inmemory

import (
	"remx/internal/config"
	"remx/internal/database"
)

var (
	// Database instance
	DB *database.Database
	// Application configuration instance
	CF *config.Configuration
)
