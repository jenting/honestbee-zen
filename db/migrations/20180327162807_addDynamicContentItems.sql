-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE dynamic_content_items (
        sn serial primary key,
        id bigint not null,
        url text not null,
        name varchar(128) not null,
        placeholder varchar(512) not null,
        default_locale_id integer not null,
        outdated boolean not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp,
        variants JSON not null
);
ALTER SEQUENCE dynamic_content_items_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX dynamic_content_items_id_unique_index ON dynamic_content_items(id);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE dynamic_content_items;
-- +goose StatementEnd
