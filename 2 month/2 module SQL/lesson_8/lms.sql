
CREATE TABLE education_level (
    education_level_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY, 
    level_name VARCHAR, -- Бакалавриат
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE education_type (
    education_type_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY, 
    type_name VARCHAR NOT NULL, -- Kunduzgi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE faculty (
    faculty_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY, 
    faculty_name VARCHAR NOT NULL, -- Iqtisodiyot
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE speciality (
    speciality_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    speciality_name VARCHAR NOT NULL, -- Bugalteriya
    description TEXT,
    price NUMERIC NOT NULL,
    faculty_id UUID NOT NULL REFERENCES faculty(faculty_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE faculty_education_level (
    faculty_id UUID NOT NULL REFERENCES faculty(faculty_id),
    education_level_id UUID NOT NULL REFERENCES education_level(education_level_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE speciality_education_type (
    speciality_id UUID NOT NULL REFERENCES speciality(speciality_id),
    education_type_id UUID NOT NULL REFERENCES education_type(education_type_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE application (
    application_ID UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR,
    middle_name VARCHAR,
    birthday DATE,
    passport_series VARCHAR(2),
    passport_number VARCHAR(7),
    passport_pinfl VARCHAR(14),
    diplom_url VARCHAR(255),
    faculty_id UUID REFERENCES faculty(faculty_id),
    education_level_id UUID NOT NULL REFERENCES education_level(education_level_id),
    speciality_id UUID REFERENCES speciality(speciality_id),
    education_type_id UUID NOT NULL REFERENCES education_type(education_type_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX application_pasport_serias_and_number_idx ON application(passport_series, passport_number);
CREATE UNIQUE INDEX application_pasport_pinfl_idx ON application(passport_pinfl);

CREATE TABLE teacher (
    teacher_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE teacher_speciality (
    teacher_id UUID NOT NULL REFERENCES teacher(teacher_id),
    speciality_id UUID REFERENCES speciality(speciality_id),
    education_level_id UUID NOT NULL REFERENCES education_level(education_level_id)
);

CREATE UNIQUE INDEX teacher_speciality_education_type_idx ON teacher_speciality( teacher_id, speciality_id, education_level_id);

CREATE TABLE course (
    course_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR NOT NULL,
    duration TIME NOT NULL,
    description TEXT,
    speciality_id UUID REFERENCES speciality(speciality_id),
    teacher_id UUID NOT NULL REFERENCES teacher(teacher_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE student (
    student_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    middle_name VARCHAR NOT NULL,
    birthday DATE NOT NULL,
    passport_series VARCHAR(2) NOT NULL,
    passport_number VARCHAR(7) NOT NULL,
    passport_pinfl VARCHAR(14) NOT NULL,
    course_number VARCHAR NOT NULL, -- 1-kurs
    course_deadline VARCHAR NOT NULL, -- 5-yil
    price NUMERIC NOT NULL, -- 12 mln
    contract_url VARCHAR(255),
    speciality_id UUID REFERENCES speciality(speciality_id),
    education_type_id UUID NOT NULL REFERENCES education_type(education_type_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
