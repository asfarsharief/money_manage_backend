package config

// Configuration - Structure for Configuration
type Configuration struct {
	Server   *ServerConfiguration
	Database *DatabaseConfiguration
}

// ServerConfiguration - Structure for ServerConfiguration
type ServerConfiguration struct {
	Env     string
	Host    string
	Port    string
	APIPath string
	WebDir  string
}

// DatabaseConfiguration - Structure for DatabaseConfiguration
type DatabaseConfiguration struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPort     string
	DBPassword string
}

// NewServerConfiguration is the constructor function to "ServerConfiguration" structure
func NewServerConfiguration(env, host string, port string, apiPath string, webDir string) *ServerConfiguration {
	return &ServerConfiguration{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
		WebDir:  webDir,
	}
}

// NewDatabaseConfiguration  is the constructor function to "DatabaseConfiguration" structure
func NewDatabaseConfiguration(dbHost, dbName, dbUser string, dbPort string, dbPassword string) *DatabaseConfiguration {
	return &DatabaseConfiguration{
		DBHost:     dbHost,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPort:     dbPort,
		DBPassword: dbPassword,
	}
}
