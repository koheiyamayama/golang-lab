CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE users (
  id integer primary key autoincrement,
  name varchar(255),
  email varchar(255) not null
);
CREATE TABLE posts (
  id integer primary key autoincrement,
  title varchar(255),
  body text,
  user_id integer not null
);
CREATE TABLE followers (
  user_id integer not null,
  follower_id integer not null
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250206141230');
