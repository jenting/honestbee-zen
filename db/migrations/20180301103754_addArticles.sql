-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
CREATE TABLE articles (
        sn serial primary key,
        section_id bigint not null,
        id bigint not null,
        author_id bigint not null,
        comments_disable boolean not null,
        draft boolean not null,
        promoted boolean not null,
        position integer not null,
        vote_sum integer not null,
        vote_count integer not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp,
        source_locale varchar(64) not null,
        outdated boolean not null, 
        outdated_locales varchar(64)[] not null,
        edited_at timestamp default localtimestamp,
        label_names varchar(128)[] not null,
        country_code varchar(64) not null
);
ALTER SEQUENCE articles_sn_seq RESTART WITH 1 INCREMENT BY 1;
CREATE UNIQUE INDEX article_id_unique_index ON articles(id);

CREATE TABLE article_translates (
        sn serial primary key,
        article_id bigint not null,
        url text not null,
        html_url text not null,
        name varchar(256) not null,
        title varchar(512) not null,
        body text not null,
        locale varchar(64) not null
);
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +goose StatementBegin
DROP TABLE article_translates;
DROP TABLE articles;
-- +goose StatementEnd