CREATE TABLE IF NOT EXISTS houses
(
  id serial primary key,
  address text not null,
  year int not null check (year >= 0),
  developer text,
  created_at timestamp not null default current_timestamp,
  update_at timestamp not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS flats
(
  id serial primary key,
  house_id int references house(id),
  number serial,
  price int not null check (price >= 0),
  rooms int not null check (rooms >= 1),
  status varchar(15) not null default 'created'
);

CREATE TABLE IF NOT EXISTS users
(
  user_id uuid primary key,
  email text not null,
  password text not null,
  user_type varchar(15) not null
);