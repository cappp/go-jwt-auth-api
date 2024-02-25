#### API de autenticação JWT feito com Go, Gin e SQLite 3

###### Como usar?

É necessário, é claro, ter o [Go](https://go.dev/) instalado, e o [curl](https://curl.se/) também (ou qualquer outro software parecido para fazer requisições http).

Baixe o repositório clonando ele pelo git ou baixando pelo GitHub.

Depois entre na pasta baixada e instale as dependências:
```sh
go mod tidy
```

Agora execute o projeto:
```sh
go run .
```

Acesse: [http://localhost:3000](http://localhost:3000) para ver a página inicial.

###### Fazendo chamadas a API

- Página inicial:
```sh
curl -c cookies.txt -b cookies.txt -X GET http://localhost:3000/
```
- Private (rota privada para usuários autentificados):
```sh
curl -c cookies.txt -b cookies.txt -X GET http://localhost:3000/private
```
- Logout (deslogar):
```sh
curl -c cookies.txt -b cookies.txt -X GET http://localhost:3000/logout
```
- Login:
```sh
curl -c cookies.txt -b cookies.txt -d '{"username": "abc", "password": "def"}' -X POST http://localhost:3000/login
```
- Signup (cadastrar):
```sh
curl -c cookies.txt -b cookies.txt -d '{"name": "x", "username": "y", "password": "z"}' -X POST http://localhost:3000/signup
```

> OBS: todas as mensagens da API estão em baianês.
