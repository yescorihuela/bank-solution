CREATE TABLE transactions(
  id VARCHAR(255) PRIMARY KEY,
  account_id VARCHAR(255) NOT NULL,
  amount DECIMAL(12,2) NOT NULL,
  kind INT NOT NULL,
  status INT NOT NULL,
  city VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW() 
);