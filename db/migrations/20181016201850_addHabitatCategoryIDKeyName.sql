
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- country sg
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360001069071','habitat','sg');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- country sg
DELETE FROM category_key WHERE category_id='360001069071' AND key_name='habitat' AND country_code = 'sg';
