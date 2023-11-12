5. Eng maximum customer kirgan film nomini chiqazing.





SELECT 
    F.title ,
    COUNT(C.customer_id)
FROM FILM AS F 
JOIN INVENTORY AS I ON I.film_id = F.film_id
JOIN RENTAL AS R ON R.inventory_id = I.inventory_id
JOIN CUSTOMER AS C ON C.customer_id = R.customer_id
GROUP BY F.title
ORDER BY COUNT DESC
limit 1
;
