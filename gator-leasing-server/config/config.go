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
}

type ServerConfig struct {
	Address string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "root",
			Password: "LKDAM6341Eastwood3!",
			Name:     "GatorLeasing",
			Charset:  "utf8",
		},
		Server: &ServerConfig{
			Address: "0.0.0.0:8080",
		},
	}
}
