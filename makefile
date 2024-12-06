run:
	go run ./cmd/project/project.go

build:
	go build -o ./bin/project ./cmd/project/project.go

test:
	FUNCTIONNEL_TEST=0 go test -cover -v ./...

functionnel:
	FUNCTIONNEL_TEST=1 go test -cover -v ./...


harbor_build:
	docker build -t go_helloworld:latest .
	docker tag go_helloworld:latest harbor.lab.rioc.fr/library/go_helloworld:latest
	docker push harbor.lab.rioc.fr/library/go_helloworld:latest