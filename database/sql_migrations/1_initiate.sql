-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL,
    name varchar(100) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE books (
    id SERIAL,
    title varchar(100) NOT NULL,
    description varchar(100) NOT NULL,
    image_url varchar(100) NOT NULL,
    release_year integer NOT NULL,
    price varchar(100) NOT NULL,
    total_page integer NOT NULL,
    thickness varchar(100) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    category_id integer NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id)
      REFERENCES category(id)
      ON DELETE CASCADE
);

-- +migrate StatementEnd