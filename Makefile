.PHONY: run test

run:
	go run cli/main.go

test:
	go test -cover -count 1 ./...
