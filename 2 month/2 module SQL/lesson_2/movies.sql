

CREATE TABLE movies (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    published_at DATE 
);

CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    description TEXT NOT NULL,
    movie_id INT NOT NULL REFERENCES movies(id),
    student_id INT NOT NULL REFERENCES students(id)
);


Relation Type

1. One  to  One
2. One  to  Many
3. Many to  Many
4. Recursive

One  To   Many
A    To   B
B    To   A


INSERT INTO movies(name, published_at) VALUES
('Shum bola', '1997-04-21'),
('Time', '2002-10-20'),
('Titanic 2', '2025-05-14'),
('Forsaj 11', null);


INSERT INTO students(id, first_name, last_name, age,  birthday) VALUES
(2, 'Bahodir', 'Isomiddinov', 22, '2001-09-28');

INSERT INTO comments(description, movie_id, student_id) VALUES
('Kutamiz', 3, 2),
('yoqmadi', 3, 2),
('serial bolib kettku', 3, 2);

('Davomi qachon chiqadi', 2, 3),
('Kommenti oqiganlar bormi', 3, 1),
('Ajoyib', 2, 2);



SELECT
    m.name,
    COUNT(c.description)
FROM
    movies AS m
LEFT JOIN comments AS c ON c.movie_id = m.id
WHERE c.student_id = 2
GROUP BY m.name
ORDER BY COUNT(c.description) DESC
-- OFFSET 1
-- LIMIT 1
;

-- COUNT()
-- SUM()
-- MAX()
-- MIN()
-- AVERAGE()
