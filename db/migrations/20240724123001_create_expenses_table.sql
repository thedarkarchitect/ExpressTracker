-- migrate:up

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    description VARCHAR(255) NOT NULL,
    amount NUMERIC(10, 2) NOT NULL, 
    date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)


-- migrate:down
DROP TABLE IF EXIST expenses;