# Go User API

Este é um projeto simples em Go utilizado para praticar a criação de uma API. A API permite cadastrar, listar, obter usuários por ID, atualizar um usuário e realizar um patch em um usuário.

## Funcionalidades

- Cadastrar usuário
- Listar usuários
- Obter usuário por ID
- Atualizar usuário
- Realizar patch em um usuário

## Requisitos

- [Docker](https://www.docker.com/get-started) instalado na máquina
- [Make](https://www.gnu.org/software/make/) instalado na máquina

## Executando o Projeto

1. Certifique-se de que o Docker está em execução na sua máquina.
2. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
3. Execute o comando make para rodar o projeto:

   ```sh
   make run

   ## Executando o Projeto

   ##Configuração

   Você pode alterar as portas utilizadas pela aplicação editando o arquivo .env ou diretamente no arquivo docker-compose.yml.

   ## Configuração

Você pode alterar as portas utilizadas pela aplicação editando o arquivo `.env` ou diretamente no arquivo `docker-compose.yml`.

### Exemplo de `.env`

  ```env
  PORT=8080
  DSN="host=postgres port=5432 user=postgres dbname=postgres password=password sslmode=disable"

### Exemplo de `docker-compose.yml`

```yaml
version: '3'
services:
  app:
    build: .
    ports:
      - "${PORT}:8080"
  db:
    image: postgres:latest
    ports:
      - "${DB_PORT}:5432"

## Documentação da API

A API possui documentação Swagger acessível em [http://localhost:{PORT}/swagger/index.html](http://localhost:{PORT}/swagger/index.html), onde `{PORT}` é a porta configurada no `.env` ou no `docker-compose.yml`.

## Endpoints da API

- **POST /users**: Cadastrar um novo usuário
- **GET /users**: Listar todos os usuários
- **GET /users/{id}**: Obter usuário por ID
- **PUT /users/{id}**: Atualizar um usuário existente
- **PATCH /users/{id}**: Realizar um patch em um usuário existente

## Contribuição

Sinta-se à vontade para abrir issues e pull requests.







   
