-- Write your migrate up statements here
ALTER TABLE users 
	ADD COLUMN role INT REFERENCES roles(id) NOT NULL DEFAULT 1;
---- create above / drop below ----
ALTER TABLE users 
	DROP COLUMN role;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
