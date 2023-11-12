CREATE USER John WITH LOGIN PASSWORD '1234';

CREATE DATABASE task_1;

CREATE TABLE teachers(
id SERIAL NOT NULL PRIMARY KEY,
first_name VARCHAR(20) NOT NULL ,
last_name CHAR(20) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE courses (
    id SERIAL NOT NULL PRIMARY KEY,
    course_name VARCHAR(20) NOT NULL,
    duration VARCHAR(10) NOT NULL,
    price NUMERIC NOT NULL,
    teacher_id INT NOT NULL REFERENCES teachers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    birthday DATE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    course_id INT NOT NULL REFERENCES courses(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO teachers(first_name,last_name)VALUES
('Asror','Samadov'),
('Abbos','Holmidov'),
('Tursin','Sattorov'),
('Shavkat','Mansurov');

INSERT INTO courses(course_name,duration,price,teacher_id)VALUES
('')

INSERT INTO
    users (first_name, last_name, age, birthday, phone_number)
VALUES
    (
        'Azim',
        'Ubaydillayev',
        25,
        '1998-05-28',
        '+99894-244-44-44'
    ),
    (
        'Ahmed',
        'Tursinov',
        18,
        '2005-04-08',
        '+99855-123-22-77'
    ),
    (
        'Tursin',
        'Zividov',
        26,
        '1997-02-23',
        '+99899-323-11-22'
    ),
    (
        'Almardon',
        'Qilichov',
        45,
        '1978-11-30',
        '+99898-454-33-77'
    ),
    (
        'Maftuna',
        'Movlonova',
        22,
        '2001-12-28',
        '+99891-678-14-44'
    ),
    (
        'Zilola',
        'Hamidova',
        19,
        '2004-05-12',
        '+99833-456-45-54'
    ),
    (
        'Nelufar',
        'Azimova',
        26,
        '1997-01-11',
        '+99875-212-33-34'
    );

CREATE TABLE courses (
    id SERIAL NOT NULL PRIMARY KEY,
    course_name VARCHAR(20) NOT NULL,
    duration VARCHAR(10) NOT NULL,
    price NUMERIC NOT NULL,
    teacher_id INT NOT NULL REFERENCES teachers (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE teachers(
id SERIAL NOT NULL PRIMARY KEY,
first_name VARCHAR(20) NOT NULL ,
last_name CHAR(20) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

INSERT INTO teachers(first_name,last_name)