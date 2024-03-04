# Minha API Simples em Go

Esta é uma API simples desenvolvida em Go que fornece funcionalidades básicas para salvar e buscar dados de usuários.

## Funcionalidades

- Salvar um novo usuário com um ID e nome.
- Buscar um usuário pelo seu ID.

## Requisitos

- Go 1.11 ou superior instalado.

## Instalação e Execução

1. Clone o repositório:

    ```bash
    git clone https://github.com/PdroBrandao/api.git
    ```

2. Navegue até o diretório do projeto:

    ```bash
    cd api
    ```

3. Execute a API:

    ```bash
    go run main.go
    ```

4. A API estará acessível em http://localhost:9999.

## Como Usar

### Salvar um novo usuário

- **Endpoint:** POST /
- **Corpo da Requisição:**

    ```json
    {
        "id": 1,
        "name": "Nome do Usuário"
    }
    ```

### Buscar um usuário

- **Endpoint:** GET /:id
- **Parâmetros de Rota:** :id - ID do usuário a ser buscado.