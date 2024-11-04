package config

type TestDBConfig struct {
	UseTestContainer bool
	UseSQLite        bool
	PostgresConfig   struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}

func GetTestDBConfig() TestDBConfig {
	// Read from env variables or return defaults
	return TestDBConfig{
		UseTestContainer: true,
		UseSQLite:        false,
	}
}
