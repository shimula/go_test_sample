create database if not exists test;

create table test.users (
    id varchar(40) primary key
) engine=innodb;

insert into test.users values ('a'),('b'),('c');