CREATE TABLE links (
    id serial PRIMARY KEY,
    shortlink text NOT NULL UNIQUE,
    url text NOT NULL
);

