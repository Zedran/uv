@ECHO OFF

go test -v ./...
go build -trimpath -ldflags "-s -w" -o build/uv.exe ./src
