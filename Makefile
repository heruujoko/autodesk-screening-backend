build:
	rm -Rf bin/
	env GOOS=linux GOARCH=amd64 go build -o bin/autodesk-identity-service-linux64 main.go
build-osx:
	rm -Rf bin/
	env GOOS=darwin GOARCH=amd64 go build -o bin/autodesk-identity-service-darwin64 main.go
build-all:
	rm -Rf bin/
	env GOOS=darwin GOARCH=amd64 go build -o bin/autodesk-identity-service-darwin64 main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/autodesk-identity-service-linux64 main.go
test:
	go test ./...