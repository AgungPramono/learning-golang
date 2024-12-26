create table customer(
    id varchar(100) not null,
    name varchar(100) not null ,
    primary key (id)
)engine=InnoDB;

DELETE FROM customer;

ALTER TABLE customer
    ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

INSERT INTO customer (id,name,email,balance,rating,birth_date,married)
VALUES ('01','agung','agung@mail.com',10000, 1.0,'1999-9-2', true);
INSERT INTO customer (id,name,email,balance,rating,birth_date,married)
VALUES ('02','joko','joko@mail.com',20000, 3.0,'1993-3-5', false);
INSERT INTO customer (id,name,email,balance,rating,birth_date,married)
VALUES ('03','budi','budi@mail.com',50000, 5.0,'1987-10-6', true);
INSERT INTO customer (id,name,email,balance,rating,birth_date,married)
VALUES ('04','angga','anggga@mail.com',70000, 4.0,'1988-11-23', true);