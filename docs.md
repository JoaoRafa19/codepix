# Codepix

## iniciando ambiente com docker

subindo os containers

```bash
docker-compose up -d
```

abrindo o shel da aplicação codepix-app-1

```bash
docker exec -it codepix-app-1 bash
```

## Testes

Executar toda a suite de testes

```bash
go test ./...
```

## Gerando o proto

```bash
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```

## Rodando o client de gRPC do evans

```bash
evans -s --path ./application/grpc/protofiles/ --path . application/grpc/protofiles/pixkey.proto  
```
