SELECT
    f.title,
    ARRAY_AGG(a.first_name||' '|| a.last_name)
FROM
    film AS f
JOIN film_actor AS fa ON f.film_id = fa.film_id
JOIN actor AS a ON a.actor_id = fa.actor_id
GROUP BY f.title
HAVING COUNT(F.title)>1
ORDER BY f.title
;