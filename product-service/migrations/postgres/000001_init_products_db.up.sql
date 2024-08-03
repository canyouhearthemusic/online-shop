CREATE TABLE IF NOT EXISTS products (
	id VARCHAR(255) PRIMARY KEY,
	title VARCHAR(255) UNIQUE NOT NULL,
	description VARCHAR(255) NOT NULL,
    price NUMERIC CHECK (price > 0) NOT NULL,
    category VARCHAR(255) NOT NULL,
    quantity NUMERIC CHECK (price >= 0) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS products_title_idx ON products(title);
CREATE INDEX IF NOT EXISTS products_category_idx ON products(category);