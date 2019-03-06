CREATE TABLE todos (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    context VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
