package config

type Config struct {
	DB     *DBConfig
	Server *ServerConfig
}

type DBConfig struct {
	Username string
	Password string
	Name     string
	Charset  string
	Address  string
}

type ServerConfig struct {
	Address string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "root",
			Password: "",
			Name:     "releasedb",
			Charset:  "utf8",
			Address:  "127.0.0.1:3306",
		},
		Server: &ServerConfig{
			Address: "0.0.0.0:8080",
		},
	}
}
