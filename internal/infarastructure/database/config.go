package database

import "os"

// Config type
type Config struct {
	DB struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

func newConfig() *Config {
	c := new(Config)
	c.DB.Host = os.Getenv("DB_HOST")
	c.DB.Port = os.Getenv("DB_PORT")
	c.DB.User = os.Getenv("DB_USER")
	c.DB.Pass = os.Getenv("DB_PASS")
	c.DB.Name = os.Getenv("DB_NAME")
	return c
}
