pragma foreign_keys = ON;
-- please also run this line after connecting to the db in the program
-- https://www.sqlite.org/foreignkeys.html

CREATE TABLE products(
  id INTEGER primary key,
  name VARCHAR(20) NOT NULL,
  price INT NOT NULL,
  description VARCHAR(50)
);

INSERT INTO products (name,price,description) VALUES ('Filter coffee',80,null),('Latte',120,'Coffee con leche'),('Oolong tea',100,null),('Ceylon tea',100,'Tea from Sri Lanka'),('Milk tea',150,'Te con leche');

CREATE TABLE cart_items(
  id INTEGER primary key,
	quantity  INTEGER NOT NULL DEFAULT 1,
	product_id INTEGER NOT NULL UNIQUE,
  inserted_at INTEGER NOT NULL DEFAULT (strftime('%s','now')), -- store as Unix timestamp
  updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
  CHECK (quantity>0)
);

-- 原本用在檢查商品是否以在購物車中，現改於在 sql 中使用 upsert ，也就是說 insert ... on conflict do update set ...
-- DO UPDATE：這會告訴 SQLite3 你想要更新已經在資料表中的列
INSERT INTO cart_items(quantity, product_id) VALUES (1,1) ON CONFLICT (product_id) DO UPDATE SET quantity = EXCLUDED.quantity+quantity, updated_at = (strftime('%s','now'));
INSERT INTO cart_items(quantity, product_id) VALUES (2,1) ON CONFLICT (product_id) DO UPDATE SET quantity = EXCLUDED.quantity+quantity, updated_at = (strftime('%s','now'));

select * from cart_items;
