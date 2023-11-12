
CREATE TABLE "Employee" (
    "id" UUID REFERENCES uuid_generate_v4(),
    "name" VARCHAR NOT NULL
);