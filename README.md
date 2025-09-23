##APClient Program

This program creates and modifies request.yaml files and submits them to an input queue
for the artificial-polyglot server.

Build and installation instructions:

To build the app for MacOS, in project root: 
```
go build -o bin/APClient .
```
to build the app for Windows, in project root:
```
GOOS=windows GOARCH=amd64 go build -o bin/APClient.exe .
Compress-Archive -Path "bin\*" -DestinationPath "apclient.zip"
```
To install the program, copy the apclient.zip file to each user's machine.  
Unzip it and the run the powerscript file `installer.ps1`.

