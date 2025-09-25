## APClient Program

This program creates and modifies request.yaml files and submits them to an input queue
for the artificial-polyglot server.

Build and installation instructions:

To download this repository:
```
git clone https://github.com/faithcomesbyhearing/APClient.git
```
To build the app for MacOS, in project root: 
```
go build -o bin/APClient .
```
to build the app for Windows, clone the project to a Windows machine, and in project root
do the instructions below. The second line create a zip file containing the executable 
and the installer.
```
go build -o bin/APClient.exe .
Compress-Archive -Path "bin\*" -DestinationPath "apclient.zip"
```
To install the program, copy the apclient.zip file to each user's machine.  
Unzip it and the run the powerscript file `installer.ps1`.

