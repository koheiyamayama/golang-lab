-- migrate:up
create table users (
  id integer,
  name varchar(255),
  email varchar(255) not null
);

create table posts (
  id integer,
  title varchar(255),
  body text,
  user_id integer not null
);

create table followers (
  user_id integer not null,
  follower_id integer not null
);
-- migrate:down
