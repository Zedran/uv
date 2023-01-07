@ECHO OFF

go test -v ./src
go build -trimpath -ldflags "-s -w" -o build/uv.exe ./src
