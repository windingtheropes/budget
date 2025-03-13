CREATE TABLE usr (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    full_name    VARCHAR(128),
    email        VARCHAR(128),
    pass         VARCHAR(128)
);

CREATE TABLE session (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    token       VARCHAR(128),
    user_id     INT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
);
