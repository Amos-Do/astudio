CREATE DATABASE a_studio_demo1;

-- enter db
\c a_studio_demo1;

CREATE TABLE IF NOT EXISTS vendor_info (
  "name" VARCHAR(50),
  intro text,
  logo_img text,
  barnner_img text,
  tel VARCHAR(50),
  open_time text,
  addr VARCHAR(225)
);