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
  house_id int references houses(id),
  price int not null check (price >= 0),
  rooms int not null check (rooms >= 1),
  status varchar(15) not null default 'created'
);

CREATE TABLE IF NOT EXISTS users
(
  id uuid primary key,
  email text unique not null,
  password text not null,
  type varchar(15) not null
);
CREATE TABLE IF NOT EXISTS subscribers
(
    id serial primary key,
    house_id int references houses(id),
    email text unique not null
);