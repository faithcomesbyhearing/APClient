APClient Program

This program creates and modifies request.yaml files and submits them to an input queue
for the artificial-polyglot server.

To build the app for MacOS, in project root: go build .

to build the app for Windows, in project root:
$env:GOOS = "windows"
$env:GOARCH = "amd64"
GOOS=windows GOARCH=amd64 go mod download
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build .
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o APClient.exe .