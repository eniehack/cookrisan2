
-- +migrate Up
CREATE TABLE ingredient (
    id uuid PRIMARY KEY,
    name TEXT,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz
);

CREATE TABLE recipe (
    id text PRIMARY KEY,
    name text NOT NULL,
    url text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz
);

CREATE TABLE crawled_ingredient (
    id uuid PRIMARY KEY,
    name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz
);

CREATE TABLE ingredient_mapping (
    id uuid PRIMARY KEY,
    ingredient_id uuid NOT NULL,
    crawled_ingredient_id uuid NOT NULL,
    FOREIGN KEY (ingredient_id) REFERENCES ingredient(id),
    FOREIGN KEY (crawled_ingredient_id) REFERENCES crawled_ingredient(id)
);

CREATE TABLE recipe_ingredient (
    id uuid PRIMARY KEY,
    recipe_id uuid,
    ingredient_id uuid,
    FOREIGN KEY ( recipe_id ) REFERENCES recipe(id),
    FOREIGN KEY ( ingredient_id ) REFERENCES crawled_ingredient(id)
);
-- +migrate Down
DROP TABLE recipe_ingredient;
DROP TABLE ingredient_mapping;
DROP TABLE crawled_ingredient;
DROP TABLE recipe;
DROP TABLE ingredient;