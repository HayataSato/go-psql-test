drop table users;


create table users (
    id         serial primary key,
    name       varchar(255) not null ,
    major      varchar(255) not null
);

