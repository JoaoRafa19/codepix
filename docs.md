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
