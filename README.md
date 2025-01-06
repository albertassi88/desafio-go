# Book API

API Rest para gerenciar livros de um usuário.

## Tecnologias
- GoLang
- MongoDB

## Endpoints

### POST /books
Cria um livro.

### GET /books
Lista todos os livros.

### GET /books/:id
Lista um livro pelo ID.

### PUT /books/:id
Atualiza um livro.

### DELETE /books/:id
Deleta um livro.

---

## Instruções para Configuração

### 1. Clone o repositório

Clone o repositório para sua máquina local:

git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio

### 2. Crie um arquivo .env

MONGO_URI=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/bookdb?retryWrites=true&w=majority

Substitua os seguintes campos:

<username>: O nome de usuário do MongoDB Atlas.
<password>: A senha do usuário do MongoDB Atlas.
<cluster>: O nome do cluster do MongoDB Atlas.

### 3. Instale as dependências

go mod tidy

### 4. Execute o projeto

go run main.go

### 5. Teste a API

O servidor estará rodando em http://localhost:8080.

Agora você pode testar os endpoints da API. 
http://localhost:8080/books. Faça um GET para obter todos os livros,
http://localhost:8080/books/<id>. Faça um GET para buscar um livro pelo ID
http://localhost:8080/books, Faça um POST para criar um livro


### Observações
Se você não possui uma conta no MongoDB Atlas, você pode criar uma gratuitamente em https://www.mongodb.com/cloud/atlas.
Certifique-se de que o banco de dados e a coleção estejam configurados corretamente no MongoDB Atlas antes de rodar a aplicação.




