CREATE TABLE coming(
    "id" UUID NOT NULL PRIMARY KEY ,
    "remaining_id" UUID NOT NULL REFERENCES "remaining"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
)
;