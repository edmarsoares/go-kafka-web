## Tecnologias utilizadas
---
- golang
- kafka
- mysql
- docker
- control-center

## Como executar o projeto local:
---

Intalando as dependencias:
- `go mod tidy`:  

Executando o projeto:
- `go run cmd/app/main.go`

## Como executar o projeto via docker-compose:
---

Na pasta raíz, rodar o seguinte comando:
- `docker-container up -d`:  

Entrar no bash do container go:
- `docker-compose exec goapp bash`

Rodar o projeto:
- `go run cmd/app/main.go`


### Criar um tópico no kafka
---

Entrar no bash do container kafka:
- `docker-compose exec kafka bash`

Criar um tópico:
- `kafka-topics --bootstrap-server=locahost:9092 --topic=products --create`

Produzir uma mensagem:

- `kafka-console-producer  --bootstrap-server=kafka:9092 --topic=products (apos o enter, informar a mensagem)`


### Criar tabela no mysql
---
Entrar no bash do container mysql:
- `docker-compose exec mysql bash`

Acessar db products:
- `mysql -uroot -p products`

Criar tabela products:
- `create table products (id varchar(255), name varchar(255), price float);`