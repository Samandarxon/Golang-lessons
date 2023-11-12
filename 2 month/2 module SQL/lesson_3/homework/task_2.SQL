-- 2. Filmlarga nechta customer kirgan

SELECT
 F.title,
 COUNT(C.customer_id) as COUNT_Customer
FROM film AS F
LEFT JOIN INVENTORY AS I ON F.film_id = I.film_id
LEFT JOIN RENTAL AS R ON R.inventory_id = I.inventory_id
LEFT JOIN CUSTOMER AS C ON C.customer_id = R.customer_id
GROUP BY
 F.title
 ORDER BY F.title
;