# Go bootstrap

## Build

This is a bootstrap goolang for starting a new project microservices using Gin.
It start 3 api in differents Port Public 8080 / Internal 8081 / Technical 8082
It use fx for injection dependency 
It ready to be connected to a BDD like mysql ( with docker compose )

## Make file

make start -> start the docker compose
make run -> run golang project
make build -> build a binarie of the project inside the folder bin
make test -> start unit test of the project 
make functionnel -> start the functionnel test of the project
make mysql -> connect to the container of mysql 