-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE sections (
        sn serial primary key,
        category_id bigint not null,
        id bigint not null,
        position integer not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp,
        source_locale varchar(64) not null,
        outdated boolean not null, 
        country_code varchar(64) not null
);
ALTER SEQUENCE sections_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX section_id_unique_index ON sections(id);

CREATE TABLE section_translates (
        sn serial primary key,
        section_id bigint not null,
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
DROP TABLE section_translates;
DROP TABLE sections;
-- +goose StatementEnd