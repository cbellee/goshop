package config

var (
	// application details
	App = struct {
		Name    string
		Usage   string
		Version string
	}{
		Name:    "productService",
		Usage:   "API for interacting with the product service",
		Version: "0.0.1",
	}

	// configuration prefix
	Prefix = "APP"

	// db migration path
	DefaultMigrationDirectory = "db/migrate"
)
