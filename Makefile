PROTO_REPO=git@github.com:eakeur-institute/comies-brunch.git

protos:
	rm -rf protos
	git clone ${PROTO_REPO} protos
	protoc --go_out=. --go-grpc_out=. -I protos protos/*/*.proto
	rm -rf protos

gen: clean protos
	go generate ./...

clean:
	find . -type f \( -name '*_mock.go' -o -name '*_mock_test.go' \) -exec rm {} +
	find . -type f \( -name '*pb.go' -o -name '*pb_test.go' \) -exec rm {} +
	rm -rf protos
	rm -rf std_protos
	rm -rf gen/*

test:
	go test ./...

setup:
	docker-compose up -d

cert:
	openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=IT Systems/OU=Eakeur/CN=localhost:50051/emailAddress=igor@eakeur.com"

	echo "CA's self-signed certificate"
	openssl x509 -in ca-cert.pem -noout -text

	# 2. Generate web server's private key and certificate signing request (CSR)
	openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Comies/OU=App/CN=localhost:50051/emailAddress=igor@eakeur.com"

	# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
	openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

	echo "Server's signed certificate"
	openssl x509 -in server-cert.pem -noout -text