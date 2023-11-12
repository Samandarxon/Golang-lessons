


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

CREATE ROLE nodirbek WITH LOGIN PASSWORD '12345';

psql -U asadbek -h localhost -d postgres

CREATE DATABASE mysql;
DROP DATABASE mysql;

DROP ROLE nodirbek;

CREATE ROLE nodirbek WITH LOGIN CREATEDB PASSWORD '12345';


