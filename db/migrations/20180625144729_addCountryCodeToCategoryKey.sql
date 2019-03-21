
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE category_key
ADD COLUMN country_code varchar(8);

-- country sg
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002338307';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002258488';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002296987';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002266307';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002289268';
UPDATE category_key SET country_code = 'sg' WHERE category_id='360000022607';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002289288';
UPDATE category_key SET country_code = 'sg' WHERE category_id='115002296947';

-- country hk
UPDATE category_key SET country_code = 'hk' WHERE category_id='115002433728';
UPDATE category_key SET country_code = 'hk' WHERE category_id='115002452107';
UPDATE category_key SET country_code = 'hk' WHERE category_id='115002452267';

-- country tw
UPDATE category_key SET country_code = 'tw' WHERE category_id='115002450907';
UPDATE category_key SET country_code = 'tw' WHERE category_id='115002432448';
UPDATE category_key SET country_code = 'tw' WHERE category_id='115002450887';

-- country jp
UPDATE category_key SET country_code = 'jp' WHERE category_id='115002442208';
UPDATE category_key SET country_code = 'jp' WHERE category_id='115002461407';
UPDATE category_key SET country_code = 'jp' WHERE category_id='115002442188';

-- country th
UPDATE category_key SET country_code = 'th' WHERE category_id='115002459987';
UPDATE category_key SET country_code = 'th' WHERE category_id='115002441308';
UPDATE category_key SET country_code = 'th' WHERE category_id='115002441288';

-- country my
UPDATE category_key SET country_code = 'my' WHERE category_id='115002440588';
UPDATE category_key SET country_code = 'my' WHERE category_id='115002438028';
UPDATE category_key SET country_code = 'my' WHERE category_id='115002456447';

-- country id
UPDATE category_key SET country_code = 'id' WHERE category_id='115002461387';
UPDATE category_key SET country_code = 'id' WHERE category_id='115002442128';

-- country ph
UPDATE category_key SET country_code = 'ph' WHERE category_id='115002440608';
UPDATE category_key SET country_code = 'ph' WHERE category_id='115002440628';
UPDATE category_key SET country_code = 'ph' WHERE category_id='115002440648';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE category_key
DROP COLUMN country_code;
