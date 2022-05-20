package config

type Config struct {
	Database Database
	Access   Access
	Server   Server
	Logger   Logger
}

type Access struct {
	ExpirationTime string
	SigningKey     string
}

type Database struct {
	User           string
	Password       string
	Host           string
	Port           string
	Name           string
	SSLMode        string
	MigrationsPath string
}

type Server struct {
	ListenPort string
	Hostname   string
}

type Logger struct {
	Environment string
}
