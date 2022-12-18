create database dev;

create table if not exists texts(
    id bigserial not null primary key,
    value text
);