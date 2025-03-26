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
CREATE TABLE entry (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    amount      INT,
    currency_id INT,
    FOREIGN KEY (user_id)
    FOREIGN KEY (user_id) REFERENCES usr(id)
);

CREATE TABLE budget_entry (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    amount      FLOAT,
    currency    VARCHAR(3),
    FOREIGN KEY (user_id) REFERENCES usr(id)
    FOREIGN KEY (currency) REFERENCES foreign_currency(id)
);
CREATE TABLE foreign_currency (
    id          VARCHAR(3) NOT NULL PRIMARY KEY,
);
CREATE TABLE tag_assignment (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_id     INT,
    entry_id    INT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
    FOREIGN KEY (entry_id) REFERENCES budget_entry(id)
);
CREATE TABLE tag (
    id          INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    tag_name    VARCHAR(128),
    user_id     INT,
    FOREIGN KEY (user_id) REFERENCES usr(id)
);