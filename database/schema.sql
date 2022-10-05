DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id varchar(20) primary key,
    password char(60)
);