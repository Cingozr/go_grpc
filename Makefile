gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

server:
	go run cmd/server/server.go

client:
	go run cmd/client/client.go
	
test:
	go test -cover -race ./...