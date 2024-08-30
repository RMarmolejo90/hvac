package initializers

import (
	"github.com/rmarmolejo90/hvac/internal/config/log"
	"github.com/rmarmolejo90/hvac/internal/db"
)

// Init initializes all the required components such as logging and database connections.
func Init() {
	// Initialize logger
	log.Init()

	// Connect to the database
	db.ConnectDB()
}
