CREATE TABLE "products"(
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" VARCHAR ,
    "barcode" VARCHAR ,
    "price" NUMERIC ,
    "category_id" uuid,
    "image_url" VARCHAR,
    "updated_at" timestamp
);