ALTER TABLE accounts ADD CONSTRAINT account_balance_check CHECK(balance >= 0);

