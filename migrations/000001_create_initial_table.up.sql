CREATE TABLE IF NOT EXISTS author(
  id serial primary key,
  email varchar not null,
  name varchar not null
);


CREATE TABLE IF NOT EXISTS post(
  id serial primary key,
  title varchar not null,
  body text,
  created_at timestamp default now(),
  author_id int not null,
  constraint fk_author_id foreign key(author_id) references author(id)
);