# CRM Go API (DDD + Hexagonal + Clean)

Projeto em Go com arquitetura em camadas (DDD + Hexagonal + Clean). Abaixo está a função de cada pasta e dos principais conceitos:

## Estrutura de pastas e responsabilidades

- `cmd/api`: ponto de entrada da aplicação. Aqui a aplicação é montada (composition root): carrega config, conecta no banco, roda migrations, cria repositórios/serviços/handlers e sobe o servidor HTTP.
- `internal/domain`: núcleo do domínio. Contém entidades e regras de negócio puras, sem dependências de infraestrutura.
  - `internal/domain/user`: entidade `User` e validações (ex.: e-mail, senha, group_id).
  - `internal/domain/user/repository`: contrato (interface) de persistência. O **Repository** define o que a aplicação precisa do banco, sem dizer como o banco é acessado.
- `internal/application`: camada de aplicação. Orquestra casos de uso e transforma dados de entrada/saída do mundo externo.
  - `internal/application/request`: DTOs de entrada usados pelos handlers (ex.: `CreateUserRequest`).
  - `internal/application/usecases`: **Use cases** representam ações do sistema (criar, listar, atualizar, deletar). Eles aplicam regras do domínio e chamam o repository via o service.
  - `internal/application/service`: agrupa dependências por contexto (ex.: `Service` com `Repo`). Serve como “porta” para os use cases, sem acoplar em infraestrutura.
- `internal/adapters`: camada de adaptação. Conecta o mundo externo (HTTP, banco, etc.) aos contratos da aplicação.
  - `internal/adapters/http/handler`: **Handlers** HTTP (controllers). Eles validam/parsam a request, chamam use cases e transformam erros em respostas HTTP.
  - `internal/adapters/http/routes`: definição e registro das rotas (mapeia paths para handlers).
  - `internal/adapters/http/shared`: helpers para respostas HTTP padronizadas.
  - `internal/adapters/postgres/user`: implementação concreta do **Repository** usando PostgreSQL (`pgx`). Contém SQL, mapeamento de erros e leitura/escrita de dados.
- `internal/infrastructure`: detalhes técnicos e bootstrap de infraestrutura.
  - `internal/infrastructure/config`: leitura de variáveis e configuração da aplicação.
  - `internal/infrastructure/database`: conexão com Postgres e execução de migrations.
  - `internal/infrastructure/database/migrations`: arquivos SQL de migração do schema.
- `internal/helpers`: utilitários compartilhados (segurança, hashing, mensagens, proteção contra SQL injection em `ORDER BY`).

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
