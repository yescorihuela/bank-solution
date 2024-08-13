CREATE TABLE customers(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  kind INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);