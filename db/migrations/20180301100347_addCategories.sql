-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE categories (
        sn serial primary key,
        id bigint not null,
        position integer not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp,
        source_locale varchar(64) not null,
        outdated boolean not null, 
        country_code varchar(64) not null
);
ALTER SEQUENCE categories_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX category_id_unique_index ON categories(id);

CREATE TABLE category_translates (
        sn serial primary key,
        category_id bigint not null,
        url text not null,
        html_url text not null,
        name varchar(256) not null,
        description text not null,
        locale varchar(64) not null
);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE category_translates;
DROP TABLE categories;
-- +goose StatementEnd