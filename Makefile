runserver:
	go run server/main.go

runclient:
	go run client/main.go -h bob

runcurl:
	curl -X POST -k https://localhost:5444/v1/example/echo -d '{"name": "charlie"}'

rungrpcurl:
	grpcurl -d '{"name": "bob"}' localhost:5443 helloworld.Greeter.SayHello

dockerbuild:
	docker build -t grpc-step-hello .

dockerrunca:
	docker run --rm --name step-ca -d -v step:/home/step -e "DOCKER_STEPCA_INIT_NAME=Smallstep" -e "DOCKER_STEPCA_INIT_DNS_NAMES=localhost,`hostname -f`,172.17.0.2,smallstep-ca" -p 9000:9000 smallstep/step-ca

dockerrunhello:
	docker run --rm --name step-hello -e "CA_FINGERPRINT=`docker exec step-ca step certificate fingerprint certs/root_ca.crt`" -e "CA_PASSWORD=`docker run -v step:/home/step smallstep/step-ca cat secrets/password`" -p 5443:5443 -p 5444:5444  grpc-step-hello 
