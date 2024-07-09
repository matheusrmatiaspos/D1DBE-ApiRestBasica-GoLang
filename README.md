# 🌐 Testes - API Rest básicas (GO Lang)

A atividade consiste na criação de duas APIs Rest (cada uma em uma linguagem de programação diferente)

## 🚧 Proposta

- Criação de duas APIs Rest (cada uma em uma linguagem de programação diferente)

- Cada API deve conter 5 rotas/endpoints, além do padrão ("/"). Cada rota deve responder às requisições GET e POST.

- Os GETs deve retornar um JSON com no mínimo 3 registros com 5 campos cada. (Ex: produtos, livros, clientes, carros, etc...)

- Os POSTs devem apenas logar no console do servidor os dados enviados pelo cliente, uma vez que ainda não estamos trabalhando com banco de dados.

- As linguagens são apenas exemplos. O(A) aluno(a) é livre para escolher a linguagem desejada.

- Usar qualquer ferramenta para testes da API (Curl, Postman, Insomnia, Insomnium, etc.). Se desejar visualizar o resultado em um navegador, utilizar alguma extensão para facilitar a visualização do JSon.

### 👉🏻 Sobre a API

- Tema: Cinema 🎬

#### ⚡ Os Endpoints
```
GET      - / 
GET/POST - /atores 
GET      - /ator/{id}
GET/POST - /diretores 
GET      - /diretor/{id}
GET/POST - /generos 
GET      - /genero/{id}
GET/POST - /filmes 
GET      - /filme/{id}
GET/POST - /analises 
GET      - /analise/{id}
```
### 🏁 Iniciando a Api
Para esta API optei por utilizar o Mux para facilitar o gerenciamento de rotas e métodos (GET e POST) sendo assim sua instalação é crucial para que a aplicação rode corretamente. Para isso utilize o comando:
```bash
go get -u github.com/gorilla/mux
```
Após isso pode executar a aplicação normalmente com o comando:
```bash
go run main.go
```

### 🎟️ Acessando EndPoints com Curl
- /
```bash
curl -X GET http://localhost:8080/

//Bem-Vindo a TestApiGo
```

- /atores [GET]
```bash
curl -X GET http://localhost:8080/atores
// [{"id":"1","name":"Morgan Freeman","birth_year":1937,"nationality":"American","movies_starred":100},{"id":"2","name":"Marlon Brando","birth_year":1924,"nationality":"American","movies_starred":50},{"id":"3","name":"Christian Bale","birth_year":1974,"nationality":"British","movies_starred":45}]
```

- /atores [POST]
```bash
curl -X POST -H 'Content-Type: application/json' -d '{"id": "4","name": "Tom Hanks","birth_year": 1956,"nationality": "Americano","movies_starred": 80}' http://localhost:8080/atores
// {"id":"4","name":"Tom Hanks","birth_year":1956,"nationality":"Americano","movies_starred":80}
```
- /ator/1
```bash
curl -X GET http://localhost:8080/ator/1
// {"id":"1","name":"Morgan Freeman","birth_year":1937,"nationality":"American","movies_starred":100}
```

### 📒 Disciplina
D1DBE - Desenvolvimento Back-End I

### 🚩 Outra API (NodeJS)
Clique [aqui](https://github.com/matheusrmatiaspos/D1DBE-ApiRestBasica-NodeJS) para acessar a outra api desenvolvida com essa mesma proposta, porém com um tema e linguagem diferentes.