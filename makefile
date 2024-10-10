run:
	go run ./cmd/project/project.go

build:
	go build -o ./bin/project ./cmd/project/project.go

test:
	FUNCTIONNEL_TEST=0 go test -cover -v ./...

functionnel:
	FUNCTIONNEL_TEST=1 go test -cover -v ./...