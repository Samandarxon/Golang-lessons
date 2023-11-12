

DO
$$
    DECLARE
        lang VARCHAR = 'English';
        obj RECORD;

    BEGIN

        FOR obj IN (
            SELECT * FROM film AS f
            JOIN language AS l ON f.language_id = l.language_id
            WHERE l.name = lang
        ) LOOP

            raise info '%', obj.title; 

        END LOOP;
    END;
$$;




CREATE OR REPLACE FUNCTION get_film_by_language(lang VARCHAR) RETURNS RECORD LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        obj RECORD;
    BEGIN

        SELECT * FROM film AS f
        INTO obj
        JOIN language AS l ON f.language_id = l.language_id
        WHERE l.name = lang;

        return obj;
    END;
$$;




CREATE OR REPLACE FUNCTION get_film_by_language(lang VARCHAR) RETURNS SETOF film LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        return QUERY (
            SELECT f.* FROM film AS f
            JOIN language AS l ON f.language_id = l.language_id
            WHERE l.name = lang
        ) ;
    END;
$$;

CREATE OR REPLACE FUNCTION get_film_by_language(lang VARCHAR) RETURNS TABLE(
    film_id INT,
    title VARCHAR
) LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        return QUERY (
            SELECT f.film_id, f.title FROM film AS f
            JOIN language AS l ON f.language_id = l.language_id
            WHERE l.name = lang
        ) ;
    END;
$$;


CREATE OR REPLACE FUNCTION get_film_by_language(lang VARCHAR) RETURNS JSONB LANGUAGE PLPGSQL
    AS
$$
    DECLARE
        objs JSONB = '[]';
        obj JSONB;
    BEGIN
        FOR obj IN (
            SELECT
                JSONB_BUILD_OBJECT(
                    'film_id', f.film_id,
                    'title', f.title,
                    'description', f.description,
                    'release_year', f.release_year,
                    'language_id', f.language_id,
                    'rental_duration', f.rental_duration,
                    'rental_rate', f.rental_rate,
                    'length', f.length,
                    'replacement_cost', f.replacement_cost,
                    'rating', f.rating,
                    'last_update', f.last_update,
                    'special_features', f.special_features,
                    'fulltext', f.fulltext
                )
            FROM film AS f
            JOIN language AS l ON f.language_id = l.language_id
            WHERE l.name = lang
        ) LOOP
            objs = objs || obj;
        END LOOP;

        return objs;
    END;
$$;



1. SELECT get_film_by_actor()
2. SELECT get_film_by_category()