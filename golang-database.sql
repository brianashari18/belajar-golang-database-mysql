show databases;


use golang_database;
CREATE TABLE customer
(
	id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDb;

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer
	ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_data DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer(id, name, email, balance, rating, birth_data, married)
VALUES
('brian','Brian', 'brian@gmail.com', 1000000, 90.0, '2003-08-18', false),
('anashari','Anashari', 'anashari@gmail.com', 500000, 85.0, '2003-09-18', true),
('puyol','Puyol', 'puyol@gmail.com', 7500000, 80.0, '2003-10-18', true);

ALTER TABLE customer
RENAME COLUMN birth_data to birth_date;

CREATE TABLE user(
	username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY(username)
) ENGINE = InnoDb;

SELECT * FROM user;

INSERT INTO user(username, password) VALUES('admin', 'admin');

CREATE TABLE comments(
	id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT,
    PRIMARY KEY(id)
) ENGINE = InnoDb;

SELECT * FROM comments;

SELECT count(*) FROM comments;
DESC comments;

create table mahasiswa(
	nim int not null auto_increment,
    nama varchar(100) not null,
    alamat varchar(100) not null,
    primary key(nim)
);

ALTER TABLE mahasiswa AUTO_INCREMENT=1301223200;

delete from mahasiswa;

select * from mahasiswa;
desc mahasiswa;