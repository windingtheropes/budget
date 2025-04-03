DROP DATABASE budget;
CREATE DATABASE budget;
USE budget;
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
    expiry      BIGINT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
);
CREATE TABLE foreign_currency (
    id          VARCHAR(3) NOT NULL PRIMARY KEY
);
CREATE TABLE transaction_entry (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    msg         VARCHAR(128),
    amount      FLOAT,
    currency    VARCHAR(3),
    unix_timestamp  BIGINT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
    -- FOREIGN KEY (currency) REFERENCES foreign_currency(id)
);
CREATE TABLE tag (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_name    VARCHAR(128),
    user_id     INT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
);
CREATE TABLE tag_assignment (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_id      INT,
    entry_id    INT,
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    FOREIGN KEY (entry_id) REFERENCES transaction_entry(id)
);

