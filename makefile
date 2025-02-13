run:
	go run ./cmd/project/project.go

build:
	go build -o ./bin/project ./cmd/project/project.go

test:
	FUNCTIONNEL_TEST=0 go test -cover -v ./...

functionnel:
	FUNCTIONNEL_TEST=1 go test -cover -v ./...


harbor_build:
	docker build -t go_api:escape_game .
	docker tag go_api:escape_game harbor.lab.rioc.fr/library/go_api:escape_game
	docker push harbor.lab.rioc.fr/library/go_api:escape_game
