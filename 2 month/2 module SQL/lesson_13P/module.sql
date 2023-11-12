CREATE DATABASE lesson_12p;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE client(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    avans INT,
    active SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP  
);


CREATE TABLE sale(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    date TIMESTAMP,
    total_amount  NUMERIC DEFAULT 0,
    paid_total_amount NUMERIC DEFAULT 0,
    remaining_total_amount NUMERIC,
    status VARCHAR DEFAULT 'new' NOT NULL,
    products_count INT DEFAULT 0,
    client_id UUID NOT NULL REFERENCES client(id),
    cash INT DEFAULT 0,
    uzcard INT DEFAULT 0,
    humo INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP 
);

 ALTER TABLE sale ALTER COLUMN  status SET DEFAULT 'new';

CREATE TABLE sale_payments(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    sale_id UUID NOT NULL REFERENCES sale(id),
    cash INT ,
    uzcard INT,
    humo INT,
    total_payment NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
     
);

CREATE TABLE sale_products(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    count INT NOT NULL,
    price INT NOT NULL,
    sum INT  ,
    sale_id UUID NOT NULL REFERENCES sale(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);