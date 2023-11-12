
CREATE TABLE "Coming_Product" (
    "Category" UUID REFERENCES "Category"("id"), 
    "ProductId" UUID REFERENCES "Product"("id"),
    "Count" INT,
    "Price" NUMERIC,
    "Total" Price NUMERIC
);