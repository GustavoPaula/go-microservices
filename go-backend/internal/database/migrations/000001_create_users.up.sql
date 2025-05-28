CREATE TABLE IF NOT EXISTS users (
  "id"          uuid  PRIMARY KEY NOT NULL  DEFAULT gen_random_uuid(),
  "name"        VARCHAR(255)  NOT NULL,
  "email"       VARCHAR(255)  NOT NULL,
  "created_at"  TIMESTAMP NOT NULL,
  "updated_at"  TIMESTAMP NOT NULL
);