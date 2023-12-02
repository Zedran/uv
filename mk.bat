@ECHO OFF

go mod tidy
go build -trimpath -ldflags "-s -w" -o build/uv.exe ./src
