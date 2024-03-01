docker run -d --name go-api-todo -p 5432:5432 -e POSTGRES_PASSWORD=go_pass postgres:13.5


se conectar com o container criado

docker exec -it go-api-todo psql -U postgres


criar database

create database go_api_todo;


criar usuario
create user go_user;

criar senhar
alter user go_user with encrypted password 'go_pass';

dar permissões
grant all privileges on database go_api_todo to go_user;




criar tabela


checar bancos disponiveis

\l

se conectar com o banco

\c go_api_todo


criar tabela
create table todos (id serial primary key, name varchar, description text, done bool);

\dt
para checar


grant all privileges on all tables in schema public to go_user;