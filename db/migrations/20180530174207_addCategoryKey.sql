
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE category_key (
        sn serial primary key,
        category_id bigint not null,
        key_name text not null,
        created_at timestamp default localtimestamp,
        updated_at timestamp default localtimestamp
);

-- country sg
INSERT INTO category_key(category_id,key_name) VALUES ('115002338307','rewards');
INSERT INTO category_key(category_id,key_name) VALUES ('115002258488','grocery');
INSERT INTO category_key(category_id,key_name) VALUES ('115002296987','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002266307','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002289268','laundry');
INSERT INTO category_key(category_id,key_name) VALUES ('360000022607','suntecConcierge');
INSERT INTO category_key(category_id,key_name) VALUES ('115002289288','tickets');
INSERT INTO category_key(category_id,key_name) VALUES ('115002296947','sumo');

-- country hk
INSERT INTO category_key(category_id,key_name) VALUES ('115002433728','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002452107','grocery');
INSERT INTO category_key(category_id,key_name) VALUES ('115002452267','myAccount');

-- country tw
INSERT INTO category_key(category_id,key_name) VALUES ('115002450907','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002432448','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002450887','grocery');

-- country jp
INSERT INTO category_key(category_id,key_name) VALUES ('115002442208','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002461407','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002442188','grocery');

-- country th
INSERT INTO category_key(category_id,key_name) VALUES ('115002459987','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002441308','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002441288','grocery');

-- country my
INSERT INTO category_key(category_id,key_name) VALUES ('115002440588','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002438028','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002456447','grocery');

-- country id
INSERT INTO category_key(category_id,key_name) VALUES ('115002461387','myAccount');
INSERT INTO category_key(category_id,key_name) VALUES ('115002442128','grocery');

-- country ph
INSERT INTO category_key(category_id,key_name) VALUES ('115002440608','grocery');
INSERT INTO category_key(category_id,key_name) VALUES ('115002440628','food');
INSERT INTO category_key(category_id,key_name) VALUES ('115002440648','myAccount');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE category_key;