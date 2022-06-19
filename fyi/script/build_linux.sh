echo "Start building for [linux]..."
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o ../fyi ../cmd/main.go
echo "Done!"