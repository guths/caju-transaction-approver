CREATE TABLE account (
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
);

CREATE TABLE merchant (
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL
);

CREATE TABLE category (
    id INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
);

CREATE TABLE mcc (
    id INT PRIMARY KEY AUTO_INCREMENT,
    mmc INT NOT NULL
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE balance (
    id INT PRIMARY KEY AUTO_INCREMENT,
    account_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL DEFAULT 0,
    category_id INT NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(ID),
    FOREIGN KEY (category_id) REFERENCES category(ID)
);

CREATE TABLE `transation` (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    account_id INT NOT NULL,
    balance_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (balance_id) REFERENCES balance(id)
);

SHOW TABLES;
