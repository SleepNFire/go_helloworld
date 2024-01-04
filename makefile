functionnel: 
	FUNCTIONNEL_TEST=1 go test -v ./...

mysql:
	docker exec -it go_bootstrap-mysql-1 mysql -u user -ppassword