


SQL - Structered Query Language
NoSQL - No Structered Query Language

PL - Programming Language

HTML - HyperText Markup Language

-- Postgresql Install
sudo apt update
sudo apt install postgresql postgresql-contrib

sudo systemctl start postgresql.service


-- Enter postgres
sudo -i -u postgres
psql


{databasename}={superadmin=#}
postgres=# 


\q      - exit psql
\du     - user list
\dt     - show tables
\!      - work terminal command
\c      - connect to database
\l      - List of databases
\d students - table Structure

CREATE ROLE nodirbek WITH LOGIN PASSWORD '12345';

psql -U asadbek -h localhost -d postgres

CREATE DATABASE mysql;
DROP DATABASE mysql;

DROP ROLE nodirbek;

CREATE ROLE nodirbek WITH LOGIN CREATEDB PASSWORD '12345';


CREATE DATABASE students;



-- TEXT
-- CHAR
-- VARCHAR(10)

-- 2000-08-10
-- 2000-08-10T22:00:30Z


-- GLOBALE VARIABLE
SELECT CURRENT_TIMESTAMP;
SELECT CURRENT_USER;
SELECT CURRENT_DATE;

CREATE TABLE students (
    id INT,
    first_name VARCHAR(34),
    last_name VARCHAR(34),
    age INT,
    birthday DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO students(id, first_name, last_name, age, birthday)
VALUES 
(1, 'Nodirbek', 'Erimbetov', 22, '2001-06-04'),
(2, 'Ilyos', 'Bektemirov', 26, '1997-06-04'),
(3, 'Zafar', 'Samadov', 18, '2005-02-15');

SELECT id, first_name, last_name, age, birthday, created_at FROM students;
SELECT * FROM students;
SELECT * FROM students WHERE first_name = 'Nodirbek';
SELECT * FROM students WHERE id = 2;

UPDATE students SET birthday = '2001-06-04' WHERE id = 1;

DELETE FROM students WHERE id = 2;

SELECT
    first_name || ' ' || last_name AS full_name,
    age,
    birthday,
    created_at
FROM
    students



SQL STATEMENTS
DDL - Data Definition Language: CREATE, DROP
DML - Data Manipulation Language: INSERT, DELETE, UPDATE
DQL - Data Query Language: SELECT
DCL - Data Control Language: GRANT, REVOKE
TCL - Transaction Control Language: BEGIN, COMMIT, ROLLBACK


