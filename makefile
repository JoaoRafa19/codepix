evans: proto
	evans --path ./application/grpc/protofiles/ --path . application/grpc/protofiles/pixkey.proto 
proto:
	protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

test:
	go test ./...

compose:
	docker-compose up -d

bash: compose
	docker exec -it codepix-app-1 bash