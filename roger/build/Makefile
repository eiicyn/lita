BINARY_NAME=roger
build:
	# use: -tags 'netcgo' to read nscd cache
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags 'netcgo' -o $(BINARY_NAME) ../cmd/main.go