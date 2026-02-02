# CRM Go API (DDD + Hexagonal + Clean)

Projeto em Go com arquitetura em camadas:

- `internal/domain`: entidades e regras de negócio puras
- `internal/application/dto/users`: DTOs de usuários
- `internal/application/service/users`: service/contexto de usuários
- `internal/application/usecases/users`: casos de uso de usuários
- `internal/adapters/http/handler/users`: handlers HTTP de usuários
- `internal/adapters/http/routes/users`: rotas HTTP de usuários
- `internal/adapters/postgres`: adaptadores de saída PostgreSQL
- `internal/infrastructure`: configuração e bootstrap técnico
- `internal/helpers`: utilitários de segurança
- `cmd/api`: ponto de entrada da aplicação

## Requisitos

- Go 1.21+
- Docker + Docker Compose

## Subir com Docker

```bash
docker compose up --build
```

API disponível em `http://localhost:8080`.

## Endpoints

- `GET /healthz`
- `POST /v1/users`
- `GET /v1/users`
- `GET /v1/users/{id}`
- `PUT /v1/users/{id}`
- `DELETE /v1/users/{id}`
- `PATCH /v1/users/{id}/group`

## Exemplo de criação de usuário

```bash
curl -X POST http://localhost:8080/v1/users \
  -H 'Content-Type: application/json' \
  -d '{
    "name":"Ana Silva",
    "email":"ana@empresa.com",
    "password":"123456",
    "telephone":"11999999999",
    "group_id": 1
  }'
```

## Exemplo de atualização de group_id

```bash
curl -X PATCH http://localhost:8080/v1/users/1/group \
  -H 'Content-Type: application/json' \
  -d '{"group_id": 2}'
```
