# Como tudo se conecta

```text
                 POSTGRESQL
              (Servidor do banco)
                      |
               simple_crud_go
                      |
              +-------+-------+
              |               |
            API Go         Beekeeper
            (db.go)        (Interface)
              |               |
            .env         (Acesso ao banco)
         (Credenciais)
```

## 1. PostgreSQL

- É um sistema de gerenciamento de banco de dados relacional (SGBD), open-source.
- É responsável por armazenar, organizar e gerenciar os dados.
- O servidor PostgreSQL precisa estar instalado e em execução na máquina que hospeda o banco.
- Possui usuários, senhas e permissões para controlar o acesso.
- Utiliza a porta `5432` por padrão.

## 2. Banco de dados

- É criado dentro do servidor PostgreSQL.
- Cada API pode ter seu próprio banco de dados.
- Exemplo: `simple_crud_go`.
- O banco armazena as tabelas e os registros utilizados pela aplicação.

## 3. Arquivo .env

- Armazena as variáveis de ambiente utilizadas pela aplicação.
- Guarda as informações necessárias para estabelecer a conexão com o banco.
- Evita deixar credenciais diretamente no código.


## 4. config/db.go

- É responsável por inicializar a conexão entre a API Go e o PostgreSQL.
- Utiliza o pacote `database/sql` para trabalhar com bancos de dados.
- Utiliza o driver `github.com/lib/pq` para comunicação com PostgreSQL.
- Utiliza `godotenv` para carregar as variáveis do arquivo `.env`.
- O método `Ping()` verifica se a conexão com o banco foi estabelecida.

## 5. Beekeeper Studio

- É um cliente gráfico para gerenciamento de bancos de dados.
- Permite visualizar tabelas, executar comandos SQL, consultar e manipular registros.
- Não é o servidor do banco e não armazena os dados.
- Conecta-se diretamente ao PostgreSQL utilizando as mesmas credenciais da aplicação.

## 6. Fluxo para iniciar uma nova API

1. Instalar o PostgreSQL na máquina (uma única vez).
2. Iniciar o serviço PostgreSQL.
3. Criar um usuário e definir sua senha.
4. Criar um banco de dados para a API.
5. Criar o arquivo `.env` com as credenciais.
6. Implementar a conexão em `config/db.go`.
7. Configurar uma conexão no Beekeeper Studio para visualizar o banco.
8. Executar a aplicação Go e testar a conexão.

**Importante:** a instalação do PostgreSQL é feita uma única vez por ambiente. Para cada nova API, normalmente você só precisa criar um novo banco, configurar as credenciais e estabelecer a conexão.