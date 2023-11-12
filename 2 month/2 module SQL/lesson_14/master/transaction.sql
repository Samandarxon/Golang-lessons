


-- TRANSACTION

TCL - Transaction Control Language

BEGIN;
BEGIN WORK;

ROLLBACK;

COMMIT;
END;


procedure <> function

function
1. transaction ishlatilmaydi
2. SELECT func_name()
3. return bor

procedure
1. transaction ishlatiladi
2. CALL getX();
3. return value qaytara olmaysiz


CREATE OR REPLACE PROCEDURE fn() LANGUAGE PLPGSQL
    AS
$$
    BEGIN

        INSERT INTO investor (first_name, phone_number, age) VALUES
        ('X', '3475487', 32);

        COMMIT;

        INSERT INTO investor (first_name, phone_number, age) VALUES
        ('Y', '2348', 23);

        ROLLBACK;
    END;
$$;


1.

CALL payment(1, 2, 10000)

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    balance NUMERIC
);

INSERT INTO users(name, balance)
VALUES
('Baxodir', 100000),
('Izzat', 50000);



