# Autodesk Identity Service Demo

this project was setup uisng Echo Framework [https://echo.labstack.com/](https://echo.labstack.com/) and supported for MySQL database

## Requirements
- Golang v1.15
- Go Modules [https://blog.golang.org/using-go-modules](https://blog.golang.org/using-go-modules)
- MySQL version 8

## Run on development mode
- `go run main.go` on the project root directory

## Build Script
Makefile is included in this repo to help on the build process. Simply run
- `make build` for linux build
- `make build-osx` for mac build
- `make build-all` for building both platform

build files shuld be available on `bin/` directory

## Config
to adjust database configuration simply edit the config.toml file
```
[database]
name="autodesk"
host="localhost"
password=""
user="root"
port=3306
```