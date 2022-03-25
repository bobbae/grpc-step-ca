runserver:
	go run server/main.go

runclient:
	go run client/main.go -h bob

runcurl:
	curl -X POST -k https://localhost:5444/v1/example/echo -d '{"name": "charlie"}'