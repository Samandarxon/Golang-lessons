-- 3. Filmlarga kirgan customer lar city chiqazing
--     films  | city_names
-- -----------+-------------
--     avatar | {Chungho, Hanoi, Rio Claro, Tashkent}
SELECT
 F.title,
 ARRAY_AGG(city) as CITES
FROM film AS F
LEFT JOIN INVENTORY AS I ON F.film_id = I.film_id
LEFT JOIN RENTAL AS R ON R.inventory_id = I.inventory_id
LEFT JOIN CUSTOMER AS C ON C.customer_id = R.customer_id
LEFT JOIN ADDRESS AS A ON A.address_id = C.address_id
LEFT JOIN CITY AS CT ON CT.city_id = A.city_id
GROUP BY
 F.title
 ORDER BY F.title
;