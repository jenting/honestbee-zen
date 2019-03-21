
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- country hk
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000148667','memberShips','hk');
-- country id
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000146188','memberShips','id');
-- country jp
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000146208','memberShips','jp');
-- country my
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000146228','memberShips','my');
-- country sg
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000148587','memberShips','sg');
-- country th
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000148907','memberShips','th');
-- country tw
INSERT INTO category_key(category_id,key_name,country_code) VALUES ('360000148927','memberShips','tw');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

-- country tw
DELETE FROM category_key WHERE category_id='360000148927' AND key_name='memberShips' AND country_code = 'tw';
-- country th
DELETE FROM category_key WHERE category_id='360000148907' AND key_name='memberShips' AND country_code = 'th';
-- country sg
DELETE FROM category_key WHERE category_id='360000148587' AND key_name='memberShips' AND country_code = 'sg';
-- country my
DELETE FROM category_key WHERE category_id='360000146228' AND key_name='memberShips' AND country_code = 'my';
-- country jp
DELETE FROM category_key WHERE category_id='360000146208' AND key_name='memberShips' AND country_code = 'jp';
-- country id
DELETE FROM category_key WHERE category_id='360000146188' AND key_name='memberShips' AND country_code = 'id';
-- country hk
DELETE FROM category_key WHERE category_id='360000148667' AND key_name='memberShips' AND country_code = 'hk';
