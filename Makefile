gen:
	protoc -I=./proto --go_out=./ --go-grpc_out=./ ./proto/*.proto

clean:
	rm ./pb/*.go

run:
	go run main.go

tidy:
	go mod tidy

test:
	go test -cover -race ./...

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080