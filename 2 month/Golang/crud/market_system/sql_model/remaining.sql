CREATE TABLE "remaining"(
    "id" UUID NOT NULL PRIMARY KEY,
    "branch_id" UUID  REFERENCES "branch"("id"),
    "product_id"  UUID  REFERENCES "product"("id") ,
    "quantity" NUMERIC ,
    "price" NUMERIC ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
);