-- stringga convertatsiya qiladi $$ 2 + 4  $$

-- SELECT $$ O'ZBEKISTON $$;


DO 
$$
    DECLARE 
        a INT = 10;
        b INT = 15;
        c INT = 20;
    
    BEGIN
        IF  a>b AND b>c THEN
        RAISE INFO 'ENG KATTA SON %',a;
        ELSIF b>a AND a>c THEN
        RAISE INFO 'ENG KATTA SON %',b;
        ELSE 
        RAISE INFO 'ENG KATTA SON %',c;
        END IF  ;
    END;

$$;