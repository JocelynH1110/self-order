select 42;
create table products(
  id integer primary key,
  name varchar(20) not null,
  price int not null,
  description varchar(50)
);

insert into products (name,price,description) values ('Filter coffee',80,null),('Latte',120,'coffee con leche'),('Oolong tea',100,null),('Cylen black tea',100,'tea from cylen'),('Milk tea',150,'te con leche');
