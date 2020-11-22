package infrastructure

import "os"

type Config struct {
	DB struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)
	c.DB.DBName = os.Getenv("DBName")
	c.DB.Host = os.Getenv("DBHost")
	c.DB.Port = os.Getenv("DBPort")
	c.DB.Username = os.Getenv("DBUsername")
	c.DB.Password = os.Getenv("DBPassword")
	c.Routing.Port = ":8080"

	return c
}
