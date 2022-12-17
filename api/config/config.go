package config

type Config struct {
	Database     Database
	Access       Access
	Server       Server
	Logger       Logger
	IDGeneration IDGeneration
}

type Access struct {
	ExpirationTime string
	SigningKey     string
}

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
	SSLMode  string
	URL      string
}

type Server struct {
	Address     string
	CORSOrigins string
}

type IDGeneration struct {
	NodeNumber string
}

type Logger struct {
	Environment string
}
