CREATE DATABASE simple_clinic;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE category(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
