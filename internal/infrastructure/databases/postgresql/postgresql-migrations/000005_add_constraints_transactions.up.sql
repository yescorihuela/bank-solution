ALTER TABLE transactions ADD CONSTRAINT fk_transactions_accounts FOREIGN KEY(account_id) REFERENCES accounts(id);