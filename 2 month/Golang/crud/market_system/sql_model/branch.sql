CREATE TABLE "branch"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR ,
    "address" VARCHAR ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
);