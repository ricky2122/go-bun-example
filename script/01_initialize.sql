-- create user table
CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(16) NOT NULL UNIQUE,
    password VARCHAR(16) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    birth_day DATE NOT NULL,
    PRIMARY KEY (id)
);

-- insert data
INSERT INTO
    users (name, password, email, birth_day)
VALUES
    (
        'user01',
        'example01',
        'example01@example.com',
        '2000-01-01'
    );

INSERT INTO
    users (name, password, email, birth_day)
VALUES
    (
        'user02',
        'example02',
        'example02@example.com',
        '2001-01-01'
    );

INSERT INTO
    users (name, password, email, birth_day)
VALUES
    (
        'user03',
        'example03',
        'example03@example.com',
        '2003-01-01'
    );