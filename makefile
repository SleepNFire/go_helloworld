start:
	docker compose up -d

run:
	go run ./cmd/stocker/stocker.go

build:
	go build -o ./bin/stocker ./cmd/stocker/stocker.go

test:
	FUNCTIONNEL_TEST=0 go test -cover -v ./...

functionnel:
	FUNCTIONNEL_TEST=1 go test -cover -v ./...

mysql:
	docker exec -it go_bootstrap-mysql-1 mysql -u user -ppassword