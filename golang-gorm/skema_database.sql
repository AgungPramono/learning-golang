create database golang_gorm;
use golang_gorm;

create table sample(
    id           varchar(100) not null ,
    name         varchar(100) not null ,
    primary key (id)
)engine =InnoDB;

select * from sample;

create table users(
    id          varchar(100)    not null ,
    password    varchar(100)    not null ,
    name        varchar(100)    not null ,
    created_at  timestamp       not null default current_timestamp,
    updated_at  timestamp       not null default current_timestamp on update current_timestamp,
    primary key (id)
)engine =InnoDB;

select * from users;

alter table users rename column name to first_name;

alter table users add column middle_name varchar(100) null after first_name;

alter table users add column last_name varchar(100) null after middle_name;


select count(*) from users where middle_name='';

SELECT * FROM `users` WHERE first_name like '%User%' AND password='1234';
SELECT * FROM `users` WHERE first_name like '%User%' OR password='12345';

SELECT * FROM `users` ORDER BY id asc, first_name desc LIMIT 5 OFFSET 5;

create table user_log(
    id          int             auto_increment,
    user_id     varchar(100)    not null ,
    action      varchar(100)    not null ,
    created_at  timestamp       not null default current_timestamp,
    updated_at  timestamp       not null default current_timestamp on update current_timestamp,
    primary key (id)
)engine =InnoDB;

select * from user_log;

alter table user_log modify created_at BIGINT NOT NULL ;
alter table user_log modify updated_at BIGINT NOT NULL ;

create table todos
(
    id              int             auto_increment,
    user_id         varchar(100)    not null ,
    title           varchar(100)    not null ,
    description     text            null,
    created_at      timestamp       not null default current_timestamp,
    updated_at      timestamp       not null default current_timestamp on update current_timestamp,
    deleted_at      timestamp       null,
    primary key (id)
)engine =InnoDB;

select * from todos;

create table wallets
(
    id          varchar(100)        not null ,
    user_id     varchar(100)        not null ,
    balance     bigint              not null ,
    created_at  timestamp           not null default current_timestamp,
    updated_at  timestamp           not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id)           references users(id)
)engine =InnoDB;

select * from wallets;
desc wallets;

create table  addresses
(
    id          bigint              not null auto_increment ,
    user_id     varchar(100)        not null ,
    address     varchar(100)        not null ,
    created_at  timestamp           not null default current_timestamp,
    updated_at  timestamp           not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id)           references users(id)
)engine=InnoDB;

desc addresses;

create table products
(
    id          varchar(100)        not null ,
    name        varchar(100)        not null ,
    price       bigint              not null ,
    created_at  timestamp           not null default current_timestamp,
    updated_at  timestamp           not null default current_timestamp on update current_timestamp,
    primary key (id)
)engine=InnoDB;

desc products;

create table user_like_product
(
    user_id     varchar(100)        not null ,
    product_id  varchar(100)        not null ,
    primary key (user_id, product_id),
    foreign key (user_id)           references users(id),
    foreign key (product_id)        references products(id)
)engine=InnoDB;

desc user_like_product;