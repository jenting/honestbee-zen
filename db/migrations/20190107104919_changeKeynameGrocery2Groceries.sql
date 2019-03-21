
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
UPDATE category_key SET key_name = 'groceries' WHERE key_name='grocery';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
UPDATE category_key SET key_name = 'grocery' WHERE key_name='groceries';
