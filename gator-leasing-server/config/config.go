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
	Migrate  bool
	Populate bool
	Clear    bool
}

type ServerConfig struct {
	Address          string
	ApiDocumentation bool
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "root",
			Password: "LKDAM6341Eastwood3!",
			Name:     "releasedb",
			Charset:  "utf8",
			Address:  "127.0.0.1:3306",
			Migrate:  true,
			Populate: true,
			Clear:    true,
		},
		Server: &ServerConfig{
			Address:          "0.0.0.0:8080",
			ApiDocumentation: true,
		},
	}
}
