
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- country ph
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000140687','memberShips','ph');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- country ph
DELETE FROM category_key WHERE category_id='360000140687' AND key_name='memberShips' AND country_code = 'ph';
