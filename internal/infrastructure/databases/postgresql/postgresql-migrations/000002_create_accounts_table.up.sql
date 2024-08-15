CREATE TABLE accounts(
  id VARCHAR(255) PRIMARY KEY,
  customer_id VARCHAR(255) NOT NULL,
  kind int NOT NULL,
  balance decimal(12,2) NOT NULL DEFAULT 0.0,
  city VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL,
  currency INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);