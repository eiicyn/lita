echo "Start building for [windows]..."
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o ../fyi.exe ../cmd/main.go
echo "Done!"
