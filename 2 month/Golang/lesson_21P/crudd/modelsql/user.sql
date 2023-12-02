CREATE TABLE users(
    "id" uuid PRIMARY KEY,
    "name" varchar(45),
    "created_at" timestamp default current_timestamp,
    "updated_at" timestamp
)
;