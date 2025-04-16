-- Write your migrate up statements here
CREATE TABLE roles (
	id SERIAL PRIMARY KEY,
	role VARCHAR(50)
);

INSERT INTO roles (role)
	VALUES ('user'), ('admin');
---- create above / drop below ----
DROP TABLE IF EXISTS roles;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
