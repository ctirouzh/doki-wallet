package config

import "os"

type Server struct {
	Host string
	Port string
}

type MySQL struct {
	Host string //DB Host
	Port string //DB Port
	Name string //DB Name
	User string //DB User
	Pass string //DB Password
}

func (mysql *MySQL) DSN() string {
	return mysql.User + ":" + mysql.Pass + "@tcp" +
		"(" + mysql.Host + ":" + mysql.Port + ")/" +
		mysql.Name + "?" + "parseTime=true&loc=Local"
}

type Config struct {
	Server Server
	MySQL  MySQL
}

func Parse() *Config {
	conf := &Config{
		Server: Server{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
		MySQL: MySQL{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
		},
	}
	return conf
}
