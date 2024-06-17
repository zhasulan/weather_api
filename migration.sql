-- auto-generated definition
create table city
(
    id      serial
        constraint city_pk
            primary key,
    name    varchar(250) not null,
    region  varchar(250),
    country varchar(250),
    lat     double precision,
    lon     double precision
);

alter table city
    owner to jappai;

