
\df  - function list
\df+ - function additinoal details

CREATE OR REPLACE FUNCTION get_x(x int, y int) RETURNS INT LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        return x + y;
    END;
$$;


CREATE OR REPLACE FUNCTION get_actor(id int) RETURNS RECORD LANGUAGE PLPGSQL 
    AS
$$
    DECLARE
        actor_data RECORD;
    BEGIN

        SELECT
            *
        FROM actor
        INTO actor_data
        WHERE actor_id = id;

        raise info 'Full Name: % %', actor_data.first_name, actor_data.last_name;
    
        return actor_data;
    END;
$$;

DO
$$
    DECLARE
        i int = 0;
    BEGIN

        LOOP
            i = i + 1;

            -- IF i > 10 THEN
            --     EXIT;
            -- END IF;

            EXIT WHEN i > 10;

            raise info 'Infinity loop .... %', i;
        END LOOP;
    END;
$$;

DO
$$
    BEGIN

        FOR i IN 1..10 BY 2 LOOP


            raise info '%', i;

        END LOOP;
    END;
$$;


DO
$$
    BEGIN

        FOR i IN REVERSE 10..1 BY 2 LOOP


            raise info '%', i;

        END LOOP;
    END;
$$;


DO
$$
    DECLARE
        i int = 0;
    BEGIN

        WHILE i < 10 LOOP

            i = i + 1;
            raise info '%', i;
        END LOOP;
    END;
$$;


SELECT reverse_number(12345678);
-- 87654321

SELECT sum_number(12345);
-- 15

