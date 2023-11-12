1. Write a Function In PostgreSql which returns most user created year
2. Write a Function in PostgreSql which raises info of all the users born in the month(June, July, August). Month should be input of function
Example: Select mostBornMonth(6) => 6 == ‘June’
3. Write a Function in PostgreSql which returns name users which were added in Year(2019). Year should be input of function
Example: Select functionname(2019)
4. Write a function in PostgreSql which gets the number of users by gender.
5. Write a function in PostgreSql which get the count of users based up on country code. Country code should be input of function.

1 - SELECT DATE_PART('Year', created_ts) from users;
2 - SELECT EXTRACT('Year' FROM created_ts) from users;

-- task_1
CREATE OR REPLACE FUNCTION get_most_user_created_year()RETURNS RECORD LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        obj RECORD;
    BEGIN
        SELECT X.dates,MAX(ARRAY_LENGTH(X.arr_year, 1))FROM(
        SELECT ARRAY_AGG(user_first_name) AS arr_year,
        EXTRACT('Year' FROM created_ts)  AS dates
        FROM users 
        GROUP BY dates
        ) AS X
        GROUP BY X.dates
        ORDER BY max DESC
        LIMIT 1
        INTO obj;
        RETURN obj;
        END;
$$;

CREATE OR REPLACE FUNCTION get_most_user_created_year()RETURNS VARCHAR LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        obj RECORD;
        datet VARCHAR;
    BEGIN
        SELECT X.dates,MAX(ARRAY_LENGTH(X.arr_year, 1))FROM(
        SELECT ARRAY_AGG(user_first_name) AS arr_year,
        EXTRACT('Year' FROM created_ts)  AS dates
        FROM users 
        GROUP BY dates
        ) AS X
        GROUP BY X.dates
        ORDER BY max DESC
        LIMIT 1
        INTO obj;
        datet=obj.dates;
        raise INFO '%',datet;
        RETURN datet;
        END;
$$;
-- task_1
-- TASK_2
CREATE OR REPLACE FUNCTION mostBornMonth(oy INT)RETURNS RECORD LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        obj RECORD;
    BEGIN
       FOR obj IN (

        SELECT ARRAY_AGG(user_first_name) AS arr_year,
        TO_CHAR(
        user_dob, 'Mon'
        ) AS mon
        FROM users 
        WHERE EXTRACT('Month' FROM user_dob) = oy
        GROUP BY mon
        ) LOOP
        END LOOP;
        
        
        raise INFO '%',obj;
        RETURN obj;
        END;
$$;

SELECT mostBornMonth(5);
-- TASK_2
-- TASK_3
CREATE OR REPLACE FUNCTION users_in_year(year INT)RETURNS RECORD LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        obj RECORD;
    BEGIN
       FOR obj IN (

        SELECT ARRAY_AGG(user_first_name) AS arr_year,
        EXTRACT('Year' FROM created_ts) AS years
        FROM users 
        WHERE EXTRACT('Year' FROM created_ts) = year
        GROUP BY years
        ) LOOP
        END LOOP;
        raise INFO '%',obj;
        RETURN obj;
        END;
$$;

SELECT users_in_year(2019); 
-- TASK_3
-- TASK_4
CREATE OR REPLACE FUNCTION users_in_year()RETURNS JSONB LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        objs JSONB = '[]';
        obj JSONB;

    BEGIN
        SELECT 
        JSONB_BUILD_OBJECT(
            'male',
            JSONB_BUILD_OBJECT(
                'count',COUNT(user_first_name),
                'users',ARRAY_AGG(user_first_name)
            )
        )
        FROM users
        INTO obj
        WHERE user_gender = 'M';
        objs = objs || obj;
     SELECT 
     JSONB_BUILD_OBJECT(
        'femail',JSONB_BUILD_OBJECT(
                'count',COUNT(user_first_name),
                'users',ARRAY_AGG(user_first_name)
            )
    )
    FROM users
    INTO obj
    WHERE user_gender = 'F';

    objs = objs || obj;
    
    RETURN objs;
    END;
$$;

SELECT users_in_year(); 
\COPY (SELECT users_in_year()) TO './users_gender.json'
-- TASK_4

-- TASK_5
CREATE OR REPLACE FUNCTION users_in_country_code(country_c INT DEFAULT 93) RETURNS  JSONB LANGUAGE PLPGSQL
AS
$$
    DECLARE 
        obj JSONB ;
    BEGIN
        FOR obj IN 
            SELECT
                JSONB_BUILD_OBJECT(
                    'users',ARRAY_AGG(user_first_name),
                    'count',COUNT(user_first_name)
                )
            FROM users
            WHERE country_code = country_c
        LOOP
        END LOOP;

    RETURN obj;
    END;

$$;

ALTER table users add column country_code  int;
update users set  country_code = 1
where user_id in (select generate_series(16,25,2))
;
-- TASK_5