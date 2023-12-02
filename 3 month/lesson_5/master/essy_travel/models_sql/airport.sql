CREATE TABLE airport(
  id UUID PRIMARY KEY,
  guid UUID,
  title VARCHAR(32),
  country_id UUID REFERENCES country(id),
  city_id UUID REFERENCES city(id),
  latitude FLOAT,
  longitude FLOAT,
  radius FLOAT,
  image VARCHAR(256),
  adress VARCHAR(256),
  timezone_id UUID,
  country VARCHAR(128),
  city VARCHAR(64),
  search_text VARCHAR(128),
  code VARCHAR(32),
  product_count INT,
  gmt VARCHAR(64),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);