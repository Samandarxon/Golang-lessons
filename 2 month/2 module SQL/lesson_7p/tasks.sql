CREATE DATABASE trade;

\c trade

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE client(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY ,
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    avans INT ,
    active SMALLINT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sale(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    date TIMESTAMP,
    total_amount NUMERIC,
    paid_total_amount  NUMERIC,
    remaining_total_amount NUMERIC,
    status  VARCHAR NOT NULL,
    products_count INT,
    client_id UUID NOT NULL REFERENCES client(id),
    cash INT,
    uzcard INT,
    humo INT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sale_payments(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    payment_type  VARCHAR NOT NULL,
    sale_id UUID NOT NULL REFERENCES sale(id),
     cash INT,
    uzcard INT,
    humo INT,
    total_payment NUMERIC NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE sale_products(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    count  INT NOT NULL,
    price INT NOT NULL,
    sum INT NOT NULL,
    sale_id UUID NOT NULL REFERENCES sale(id),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);