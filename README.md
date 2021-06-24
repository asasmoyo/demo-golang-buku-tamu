Demo Golang Buku Tamu
=====================

Db setup:

```
create database tamudb;
\c tamudb;
create table tamus (id serial4 primary key, name varchar(200) not null, keperluan text not null, tanggal timestamptz default now());
```