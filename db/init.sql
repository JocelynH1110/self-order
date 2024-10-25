pragma foreign_keys = ON;
-- please also run this line after connecting to the db in the program
-- https://www.sqlite.org/foreignkeys.html

create table products(
  id integer primary key,
  name varchar(20) not null,
  price int not null,
  description varchar(50)
);

insert into products (name,price,description) values ('Filter coffee',80,null),('Latte',120,'Coffee con leche'),('Oolong tea',100,null),('Ceylon tea',100,'Tea from Sri Lanka'),('Milk tea',150,'Te con leche');

create table cart_items(
  id integer primary key,
	quantity  integer not null ,
	product_id integer ,
  inserted_at integer , -- store as Unix timestamp
  updated_at integer ,
  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
);

insert into cart_items(quantity,inserted_at) values (?,strftime('%s','now'));
