run:
	go run cmd/main.go

build:
	cd cmd && go build -o ../bin/main

run-build:
	./bin/main