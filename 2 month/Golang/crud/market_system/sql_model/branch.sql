CREATE TABLE "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR ,
    "address" NUMERIC ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
);