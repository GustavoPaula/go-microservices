CREATE TABLE IF NOT EXISTS users (
  "id"          uuid  PRIMARY KEY NOT NULL  DEFAULT gen_random_uuid(),
  "name"        VARCHAR(255)  NOT NULL,
  "email"       VARCHAR(255)  NOT NULL,
  "password"    VARCHAR(255)  NOT NULL,
  "is_active"   BOOLEAN DEFAULT TRUE,
  "created_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at"  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "deleted_at"  TIMESTAMP
);

INSERT INTO users (id, name, email, password, is_active, created_at, updated_at)
VALUES (
  gen_random_uuid(), 
  'admin', 
  'admin@admin.com', 
  '$2a$14$R704Gt1pYcL2X/8KAqqsvO8HnGm13DVo3HILfcZSxavsUxIQG6mo2', 
  TRUE, 
  current_timestamp,
  current_timestamp
  );