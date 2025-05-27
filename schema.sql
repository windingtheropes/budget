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
CREATE TABLE transaction_type (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    type_name   VARCHAR(128)
);
CREATE TABLE transaction_entry (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    type_id     INT,
    msg         VARCHAR(128),
    amount      FLOAT,
    currency    VARCHAR(3),
    unix_timestamp  BIGINT,
    vendor      VARCHAR(128),
    FOREIGN KEY (user_id) REFERENCES usr(id),
    FOREIGN KEY (type_id) REFERENCES transaction_type(id)
);
CREATE TABLE tag (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_name    VARCHAR(128)
);
CREATE TABLE tag_ownership (
    id  INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_id  INT NOT NULL, 
    user_id INT NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag(id)
);
CREATE TABLE tag_assignment (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_id      INT NOT NULL,
    entry_id    INT NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    FOREIGN KEY (entry_id) REFERENCES transaction_entry(id)
);

-- new stuff

CREATE TABLE budget (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    type_id     INT,
    budget_name VARCHAR(128),
    goal         FLOAT,
    FOREIGN KEY (user_id) REFERENCES usr(id),
    FOREIGN KEY (type_id) REFERENCES transaction_type(id)
);

CREATE TABLE budget_entry (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    transaction_id  INT,
    budget_id   INT,
    amount      FLOAT,
    FOREIGN KEY (transaction_id) REFERENCES transaction_entry(id),
    FOREIGN KEY (budget_id) REFERENCES budget(id)
);

CREATE TABLE tag_budget (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_id      INT,
    budget_id   INT,
    goal        FLOAT,
    type_id     INT,
    FOREIGN KEY (tag_id) REFERENCES tag(id),
    FOREIGN KEY (budget_id) REFERENCES budget(id),
    FOREIGN KEY (type_id) REFERENCES transaction_type(id)
);