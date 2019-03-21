
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE articles ADD COLUMN click_count bigint default 0;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE articles DROP COLUMN click_count;
