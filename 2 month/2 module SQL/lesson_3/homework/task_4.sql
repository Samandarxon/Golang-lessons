4. payment_date pul tolangan kinolar nomini chiqazing va umumiy summa
Masalan:
payment_date | film_names          | total_amount
-------------+---------------------+--------------
  2007-02-16 | {avatar, jon uik3 } | 220.23


-- task 4
SELECT
 TO_CHAR(P.payment_date, 'yyyy-mm-dd') as dates,
 ARRAY_AGG(F.title) AS film_names,
 SUM(P.amount) AS total_amount
FROM PAYMENT AS P
JOIN RENTAL AS R ON R.rental_id = P.rental_id
JOIN INVENTORY AS I ON I.inventory_id = R.inventory_id
JOIN FILM AS F ON F.film_id = I.film_id
GROUP BY 
dates
ORDER by dates
;