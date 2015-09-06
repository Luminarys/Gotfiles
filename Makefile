all:
	go fmt src/*.go
	go vet src/*.go
	go build -o gf src/*.go
install:
	cp gf /usr/local/bin/
