@ECHO OFF

go build -trimpath -ldflags "-s -w" -o build/uv.exe ./src
