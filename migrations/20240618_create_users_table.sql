create TABLE users(
    id int auto_increment primary key,
    name VARCHAR(30) not null,
    username varchar(30) not null,
    email varchar(256) not null
)