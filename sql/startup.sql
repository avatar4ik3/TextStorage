-- create database dev;
create table if not exists texts(
    id bigserial not null primary key,
    description varchar,
    value text
);

create table if not exists groups(
    id bigserial not null primary key,
    description varchar
);

create table if not exists relations(
    text_id bigserial not null REFERENCES texts(id),
    group_id bigserial not null REFERENCES groups(id),
    CONSTRAINT PK_Relation PRIMARY KEY (text_id, group_id)
);