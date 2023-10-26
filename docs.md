# Codepix

## iniciando ambiente com docker

subindo os containers

```bash
docker-compose up -d
```

para abrir o shell de qualquer container

```bash
docker exec -it <container_name> bash
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

## Cobra CLI

``` bash
cobra-cli init 
```

add grpc

``` bash
cobra-cli add grpc
```

### Inicia o gRPC server usando o cobra

``` bash
go run main.go grpc
```


## Kafka

### Listando topicos

``` bash
kafka-topics --list --bootstrap-server=localhost:9092
```

### Consumindo mensagens de um topico

``` bash
kafka-console-consumer --topic=teste --bootstrap-server=kafka:9092
```