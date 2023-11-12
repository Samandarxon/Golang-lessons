CREATE DATABASE magazin;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "uuid-ossp";

CREATE TABLE category (
    id SERIAL PRIMARY KEY NOT NULL,
    category_name varchar(40) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE shop (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR NOT NULL,
    location VARCHAR,
    phone VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE staff(
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    shop_id INT NOT NULL REFERENCES shop(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product(
    id SERIAL PRIMARY KEY ,
    product_name VARCHAR NOT NULL,
    product_price DECIMAL  NOT NULL,
    shop_id INT NOT NULL REFERENCES shop(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product_category(
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES product(id),
    category_id INT NOT NULL REFERENCES category(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

SELECT DISTINCT
SH.name,
ARRAY_AGG(SF.first_name || '  '||SF.last_name )
FROM staff AS SF
JOIN shop AS SH ON SH.id = SF.shop_id
GROUP BY SH.name

;