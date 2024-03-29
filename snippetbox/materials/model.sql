--PostgreSQL 16
CREATE TABLE cf_menu (
id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
pos_name varchar not null,
price numeric not null default 1);

INSERT INTO cf_menu (pos_name, price) VALUES ('espresso coffee', 80);
INSERT INTO cf_menu (pos_name, price) VALUES ('latte coffee', 110);
INSERT INTO cf_menu (pos_name, price) VALUES ('cappuccino coffee', 120);
INSERT INTO cf_menu (pos_name, price) VALUES ('americano coffee', 90);
INSERT INTO cf_menu (pos_name, price) VALUES ('chocolate cake', 80);
INSERT INTO cf_menu (pos_name, price) VALUES ('blueberry cake', 90);
INSERT INTO cf_menu (pos_name, price) VALUES ('apple tart', 100);

CREATE TABLE cf_person (
id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY ,
name varchar not null,
age integer not null default 10,
gender varchar default 'female' not null ,
address varchar);

INSERT INTO cf_person (name, age, gender, address) VALUES('Andrey', 21, 'male', 'Moscow');

CREATE TABLE cf_orders (
id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY ,
surname varchar not null,
name varchar not null,
amount numeric not null,
order_date timestamp not null default current_timestamp);

