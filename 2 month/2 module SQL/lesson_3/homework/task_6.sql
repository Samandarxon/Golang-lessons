
6. Biron bir country ga tegishli customer lar royxatini chiqazing

WHERE country = 'Japan'

    country | customer_name
------------+-------------
      Japan | Patricia
      Japan | Elizabeth
      Japan | Carol
      Japan | Ruth
      Japan | Sharon
      Japan | Michelle
      Japan | Laura
      Japan | Sarah
      Japan | Kimberly
      Japan | Deborah



SELECT 
C.country,
CU.first_name ||' '||CU.last_name as customer_name
FROM COUNTRY AS C
JOIN CITY AS CT ON C.country_id = CT.country_id
JOIN ADDRESS AS A ON A.city_id = CT.city_id
JOIN CUSTOMER as CU ON CU.address_id = A.address_id
WHERE C.country = 'Japan'
;