# Soccer Services

This project soccer API that give a list of soccer players in a team.

## Getting Started

See api documentations [here](https://github.com/Dandi-Pangestu/efishery-service/blob/master/API.md).

### Prerequisites

* [Go](https://golang.org/)
* [MySQL](https://www.mysql.com/)

### Installing (Fetching Service)

This step for installing and run the service

Prepare config
```
cp shared/config/config.default shared/config/config.yml
```

Prepare database
```
Make sure to create database team-player in your instance and dont forget to change instance, username, and password in config.yml

driver: mysql
instance: localhost
port: 3306
dbname: team-player
username: root
password:
options: charset=utf8&parseTime=True&loc=UTC
```

Install dependencies
```
go mod tidy
```

Running server
```
go build cmd/service-team-player/main.go
./main
```

The server will running at localhost:8080

### Testing in Fetching Service

Running test
```
go test -v ./...
```
