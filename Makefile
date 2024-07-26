nm:
	go vet && golangci-lint run src && nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

tst:
	go test ./src/... -bench=. -v -benchmem -cover

test-report:
	go test -v -cover -json  ./src/linkedlist -bench . -benchmem | go-test-html-report 