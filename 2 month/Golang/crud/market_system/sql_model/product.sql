CREATE TABLE "products"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR ,
    "price" NUMERIC ,
    "category_id" UUID NOT NULL REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
);