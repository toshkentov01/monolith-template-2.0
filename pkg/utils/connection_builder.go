package utils

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/toshkentov01/template/config"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	config := config.Config()
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDB,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			":%d",
			config.ServerPort,
		)

	case "migration":
		// URL for Migration
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresDB,
		)

	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
