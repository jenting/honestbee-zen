
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360001339052','bungkus','my');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM category_key WHERE category_id='360001339052' AND key_name='bungkus' AND country_code = 'my';
