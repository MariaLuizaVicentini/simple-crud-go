# MEU PRIMEIRO CRUD COM A LINGUAGEM GO

Esse repositório documenta meu primeiro contato com a linguagem. 

Golang é uma linguagem de programação criada pelo Google que é conhecida por sua eficiência, alto desempenho e suporte nativo a concorrência. Muitas empresas de tecnologia como Docker, Kurbenetes e Uber utilizam Go.

A ideia é construir sem IA uma API REST em Go com todas as operações do CRUD:
- READ
- DELETE
- GET
- UPDATE

Esse é o vídeo que usei como referência: https://www.youtube.com/live/rQnhcO5Q68A?si=YW_8dwPbTfLjqpWi

---

A primeira coisa a ser fazer é criar o módulo Go
- gerencia as dependencias da app
- semelhante ao package.json

```
go mod init nome-modulo
```

---

# Dependencias

- Beekeeper Studio é um gerenciador de banco de dados e interface gráfica para SQL, vou utiliza-lo no projeto ao invés do postico.

- rotas http
```
go get -u github.com/gorilla/mux
```

- conexão com o postgres
```
go get -u github.com/lib/pq
```

- variaveis de ambiente
```
go get -u github.com/joho/godotenv
```

