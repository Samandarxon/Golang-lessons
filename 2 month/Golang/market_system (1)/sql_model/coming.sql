
CREATE TABLE "Coming"(
    "Datetime" TIMESTAMP NO NULL,
    "Employee" UUID REFERENCES "Employee"("id"),
    "Address" VARCHAR,
    "Status" VARCHAR NOT NULL,        --> in_proccess, success
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updatet_at" TIMESTAMP
);