.PHONY: client fetcher server 

all:
	go build -mod=mod -o rs-client client/main.go
	go build -mod=mod -o rs-fetcher fetcher/main.go
	go build -mod=mod -o rs-server server/main.go

linux:
	GOARCH=amd64 GOOS=linux go build -mod=mod -o rs-client client/main.go
	GOARCH=amd64 GOOS=linux go build -mod=mod -o rs-fetcher fetcher/main.go
	GOARCH=amd64 GOOS=linux go build -mod=mod -o rs-server server/main.go

client:
	go build -mod=mod -o rs-client client/main.go

fetcher:
	go build -mod=mod -o rs-fetcher fetcher/main.go

server:
	go build -mod=mod -o rs-server server/main.go

runclient:
	go run client/main.go;

runfetcher:
	go run fetcher/main.go;

runserver:
	go run server/main.go;
