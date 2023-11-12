-- 1. Filmlarga qancha umumiy pul tushgan
-- SUM() foydalansiz

SELECT
 F.title,
  SUM(P.amount)
FROM film AS F
LEFT JOIN INVENTORY AS I ON F.film_id = I.film_id
LEFT JOIN RENTAL AS R ON R.inventory_id = I.inventory_id
LEFT JOIN PAYMENT AS P ON P.rental_id = R.rental_id
GROUP BY
 F.title
 ORDER BY F.title
--  P.amount 
;