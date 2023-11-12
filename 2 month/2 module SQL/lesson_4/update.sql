UPDATE film SET title=title ||'   ' ||(SELECT name FROM category WHERE category_id = 2)
WHERE film_id IN (SELECT film_id FROM film_category WHERE category_id = 2)
;

SELECT title FROM film WHERE film_id IN (SELECT film_id FROM film_category WHERE category_id = 2);