SELECT DISTINCT CITY FROM STATION
WHERE LOWER(CITY) RLIKE '(^[^aeiou])|([^aeiou]$)';