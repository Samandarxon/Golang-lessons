SELECT DISTINCT CITY
FROM STATION
WHERE
LOWER(city) RLIKE '^[aeiou].*[aeiou]$';
