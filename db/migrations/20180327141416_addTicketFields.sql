-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE ticket_fields (
        sn serial primary key,
        id bigint not null,
        url text not null,
        type varchar(64) not null,
        title varchar(128) not null,
        raw_title varchar(128) not null,
        description varchar(512) not null,
        raw_description varchar(512) not null,
        position integer not null,
        active boolean not null,
        required boolean not null,
        collapsed_for_agents boolean not null,
        regexp_for_validation varchar(256) not null,
        title_in_portal varchar(256) not null,
        raw_title_in_portal varchar(256) not null,
        visible_in_portal boolean not null,
        editable_in_portal boolean not null,
        required_in_portal boolean not null,
        tag varchar(64) not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp,
        removable boolean not null,
        custom_field_options JSON not null,
        system_field_options JSON not null
);
ALTER SEQUENCE ticket_fields_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX ticket_fields_id_unique_index ON ticket_fields(id);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE ticket_fields;
-- +goose StatementEnd
