# GatorLeasing
GatorLeasing is a website to facilitate subleasing agreements between Gainesville residents. Users can post subleasing opportunities and communicate with interested individuals.

Frontend:

- Milena Petrovic
- Lily Penoyer  W

Backend:

- Nick Rice
- Jacob Immich

## Installation
First clone the repository
```bash
git clone https://github.com/milenapetrov/GatorLeasing
```

# Go Backend API

## Navigate to Server Directory

> **Step 1** - Once the project is downloaded, change the directory to `gator-leasing-server`

```bash
$ cd gator-leasing-server
```

## Install Go Dependencies

> **Step 2** - Install all dependencies

```bash
$ go get ./...
```

## Database Configuration

> **Step 3** - Update the config file with your MYSQL password

Before running the server, you should set the database config with your values in [config.go](https://github.com/milenapetrov/GatorLeasing/blob/main/gator-leasing-server/config/config.go)
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "root",
			Password: "YOUR_PASSWORD",
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
```

## Run Backend

> **Step 4** - Start the server on http://localhost:3000

```bash
go run main.go
```

The backend api will now be running on http://localhost:3000

# Angular Frontend (use another terminal)

## Naivate to Client Directory

> **Step 1** - Once the project is downloaded, change the directory to `gator-leasing-client`

```bash
$ cd gator-leasing-client
```

## Install Angular Dependencies

> **Step 2** - Install dependencies via NPM

```bash
$ npm i
```

## Run Frontend

> **Step 3** - Start the frontend on http://localhost:4200

```bash
$ npm start
```

The frontend client will now be running on http://localhost:4200