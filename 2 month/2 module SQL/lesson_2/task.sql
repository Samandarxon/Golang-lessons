-- 1. Studentlar qaysi kinolarga comment yozgan
-- Bahodir | {Titanic, Forsaj, Time}
-- Nodirbek | {Titanic}
CREATE DATABASE movies;

CREATE TABLE students(
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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

INSERT INTO
    movies(name, published_at)
VALUES
    ('Shum bola', '1997-04-21'),
    ('Time', '2002-10-20'),
    ('Titanic 2', '2025-05-14'),
    ('Forsaj 11', null);

INSERT INTO
    students(first_name, last_name, age, birthday)
VALUES
    ('Bahodir', 'Isomiddinov', 22, '2001-09-28'),
    ('Islom', 'Sattorov', 23, '2000-12-28'),
    ('Azamat', 'Shodmonov', 22, '2001-09-28'),
    ('Atham', 'Mamatov', 22, '2001-09-28');

INSERT INTO
    comments(description, movie_id, student_id)
VALUES
    ('Kutamiz', 3, 2),
    ('yoqmadi', 3, 2),
    ('serial bolib kettku', 3, 4),
    ('Davomi qachon chiqadi', 2, 3),
    ('Kommenti oqiganlar bormi', 3, 1),
    ('Ajoyib', 2, 2);


-- UPDATE
--     comments AS c
-- SET
--     -- postgres FTW
--     movie_id = u2.movie_id,
--     student_id = u2.student_id
-- from
--     (
--         values
--             (1, 3, 2),
--             (2, 3, 2),
--             (3, 3, 4),
--             (4, 2, 3),
--             (5, 1, 1),
--             (6, 2, 2)
--     ) as u2(id, movie_id, student_id)
-- where
--     u2.id = c.id;

-- SELECT * FROM
    SELECT
        s.first_name || ' ' || s.last_name AS fullname,
        ARRAY_AGG(m.name) AS movies,
        COUNT (m.name) AS description_count
    FROM
        students AS s
        LEFT JOIN comments AS c ON c.student_id = s.id
        LEFT JOIN movies AS m ON c.movie_id = m.id AND c.student_id = s.id
    GROUP BY
        fullname
    ORDER BY
        COUNT(c.description) DESC;