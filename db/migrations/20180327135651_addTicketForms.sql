-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE ticket_forms (
        sn serial primary key,
        id bigint not null,
        url text not null,
        name varchar(128) not null,
        raw_name varchar(128) not null,
        display_name varchar(256) not null,
        raw_display_name varchar(256) not null,
        end_user_visible boolean not null,
        position integer not null,
        active boolean not null,
        in_all_brands boolean not null,
        restricted_brand_ids bigint[] not null,
        ticket_field_ids bigint[] not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp
);
ALTER SEQUENCE ticket_forms_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX ticket_forms_id_unique_index ON ticket_forms(id);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE ticket_forms;
-- +goose StatementEnd
