nm:
	go vet && golangci-lint run src && nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go