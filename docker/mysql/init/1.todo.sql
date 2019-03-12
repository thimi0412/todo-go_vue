CREATE TABLE todos (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    context VARCHAR(255) NOT NULL,
    limit_date DATETIME NOT NULL,
    insert_date DATETIME NOT NULL,
    updated_date DATETIME NOT NULL,
    PRIMARY KEY (id)
);
