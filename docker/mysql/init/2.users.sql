CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(32) NOT NULL,
    password VARCHAR(32) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (password)
);
