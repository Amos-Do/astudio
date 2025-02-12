-- enter db
\c a_studio;

-- filling of users
INSERT INTO users (name, birthday, account, password)
  VALUES ('Tom', '1994-11-29', 'tom1129@gmail.com', '1234');

-- load the pgcrypto extension to gen_random_uuid ()
CREATE EXTENSION pgcrypto;

-- filling of vendors
INSERT INTO vendors (id, vendor_uuid, db_name)
  VALUES (DEFAULT, gen_random_uuid(), 'a_studio_demo1'),
  (DEFAULT, gen_random_uuid(), 'a_studio_test1');

-- filling of user_vendors
INSERT INTO user_vendors (user_id, vendor_id)
  VALUES (1, 1),
  (1, 2);

-- filling of customer
INSERT INTO customer (name, birthday, account, line_id)
  VALUES ('Nick', '1984-10-10', 'nick1010@gmail.com', gen_random_uuid()),
  ('May', '1982-6-5', 'may0605@gmail.com', gen_random_uuid());

-- filling of customer_members
INSERT INTO customer_members (customer_id, vendor_id)
  VALUES (1, 2),
  (1, 1),
  (2, 1);