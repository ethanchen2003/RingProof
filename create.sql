drop table if exists userReport, spamPhone, users;

create table IF NOT EXISTS users (
        userId int not null AUTO_INCREMENT,
        lastName varchar(255) not null,
        firstName varchar(255) not null,
        bod date not null,
        gender char(1) not null,
        phone varchar(10) not null,
        email varchar(50) not null,
        username varchar(50) not null,
        passwrod varchar(50) not null,
        primary key (userId),
        UNIQUE (username),
        CONSTRAINT check_gender CHECK (gender in ('m', 'f'))
);

create table IF NOT EXISTS spamPhone (
        phone varchar(10) not null primary key,
        reportNumber int DEFAULT 0
);

create table IF NOT EXISTS userReport (
        id int not null AUTO_INCREMENT,
        reporter int not null,
        phone varchar(10) null,
        reportDateTime DATETIME DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        FOREIGN KEY (reporter) REFERENCES users(userId),
        FOREIGN KEY (phone) REFERENCES spamPhone(phone),
        UNIQUE (reporter, phone)
);
