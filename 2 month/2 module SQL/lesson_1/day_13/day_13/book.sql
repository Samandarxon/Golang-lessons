

CREATE TABLE author (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);

CREATE TABLE book (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    price NUMERIC NOT NULL,
    author_id INT NOT NULL REFERENCES author(id)
);


INSERT INTO author(name) VALUES
('Abdullo Qodiriy'),
('Cho''lpon'),
('Navoiy');

truncate TABLE author restart identity;

INSERT INTO book(name, price, author_id) VALUES
('O''tkan kunlar', 25000, 1);

INSERT INTO book(name, price, author_id) VALUES
('Ko''klam ruhi', 40000, 2),
('Hamsa', 100000, 3),
('Farhod va shirin', 340000, 3);


SELECT
    b.id,
    b.name AS book_name,
    b.price,
    a.name AS author_name
FROM
    book AS b
JOIN author AS a ON b.author_id = a.id
;


Task 1

users table
    --name
    --age
    --phone_number

course table
    --name
    --duration
    --price

teacher table
    --name

     user_name      | teacher_name |  course_name  |         course_price  | course_duration         
--------------------+-----+------------+----------------------------
 Zafar Samadov      |  ...         |        ...    |         ....          |
 Nodirbek Erimbetov |  ...         |        ...    |         ....          |



