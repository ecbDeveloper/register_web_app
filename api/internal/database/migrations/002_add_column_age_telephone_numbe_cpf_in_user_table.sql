-- Write your migrate up statements here
ALTER TABLE users 
	ADD COLUMN cpf VARCHAR(11),
	ADD COLUMN phone_number VARCHAR(20),
	ADD COLUMN age INT;
---- create above / drop below ----
ALTER TABLE users 
	DROP COLUMN cpf,
	DROP COLUMN	phone_number,
	DROP COLUMN	age;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
