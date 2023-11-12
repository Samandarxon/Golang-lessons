SELECT t.count 
FROM generate_series(1,1000) AS t(count);


INSERT INTO student (id,first_name) VALUES
SELECT t.number, 'USER' || t.number 
FROM generate_series(1,100000000000000000) AS t(number);