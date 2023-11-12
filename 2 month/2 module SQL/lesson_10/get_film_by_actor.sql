CREATE OR REPLACE FUNCTION get_film_by_actor(actor_name VARCHAR) RETURNS JSONB LANGUAGE PLPGSQL
    AS

$$
    DECLARE 
    objs JSONB = '[]';
    obj JSONB ;

    BEGIN
            
        --  raise info '%',actor_name;
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
            JOIN film_actor AS fac ON f.film_id = fac.film_id
            JOIN actor AS ac on ac.actor_id = fac.actor_id
            WHERE ac.first_name = actor_name
        ) LOOP
        --  raise info '%',objs;
         objs = objs || obj;

         END LOOP;
         return objs;
    END;

$$;

SELECT get_film_by_actor('Jennifer');

-- write postgres from to csv file  
 \COPY (SELECT get_film_by_actor('Nick') )TO './actor.csv' WITH DELIMITER ',' CSV HEADER;

-- write postgres from to json file  
\COPY (SELECT get_film_by_actor('Vivien') ) TO './actor.json';
