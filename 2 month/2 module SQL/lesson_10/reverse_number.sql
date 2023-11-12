CREATE OR REPLACE FUNCTION reverse_number(X INT) RETURNS INT LANGUAGE PLPGSQL
    AS
$$  
    DECLARE 
        D INT = 0;
    BEGIN
        LOOP
            D = D *10+ X%10; 
    -- RAISE INFO '%',X;
            X = X / 10;
    RAISE INFO '%',X;
            EXIT WHEN X < 10;
        
        END LOOP;

        RAISE INFO '%',D;

        RETURN D;
    END;    
$$;

-- SELECT reverse_number(12345678);
