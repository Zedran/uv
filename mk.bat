@ECHO OFF

go mod tidy
go test ./...
go build -trimpath -ldflags "-s -w" -o build/uv.exe ./src
