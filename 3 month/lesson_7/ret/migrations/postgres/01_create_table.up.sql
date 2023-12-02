CREATE TABLE countries (
    guid VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255),
    code VARCHAR(2),
    continent VARCHAR(2)
);


CREATE TABLE cities (
    guid VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255),
    country_id VARCHAR(36) REFERENCES countries(guid) ON DELETE CASCADE,
    city_code VARCHAR(255),
    latitude VARCHAR(255),
    longitude VARCHAR(255),
    "offset" VARCHAR(255),
    timezone_id VARCHAR(36),
    country_name VARCHAR(255)
);

CREATE TABLE buildings (
    guid VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255),
    country_id VARCHAR(36) REFERENCES countries(guid) ON DELETE CASCADE,
    city_id VARCHAR(36) REFERENCES cities(guid) ON DELETE CASCADE,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    radius DOUBLE PRECISION,
    image VARCHAR(255),
    address VARCHAR(255),
    timezone_id VARCHAR(36),
    country VARCHAR(255),
    city VARCHAR(255),
    search_text VARCHAR(255),
    code VARCHAR(10),
    product_count INTEGER,
    gmt VARCHAR(19)
);
