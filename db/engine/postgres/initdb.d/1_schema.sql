set client_encoding = 'UTF8';

-- ---- create ----
-- CREATE TABLE bans (
--     id     integer PRIMARY KEY,
--     name   VARCHAR(20),
--     stock  integer
-- );


-- ---- create ----
-- CREATE TABLE patties (
--     id     integer PRIMARY KEY,
--     name   VARCHAR(20),
--     stock  integer
-- );


---- create ----
CREATE TABLE vegetables (
    id     integer PRIMARY KEY,
    name   VARCHAR(20),
    stock  integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


---- create ----
CREATE TABLE ingredients (
    id     integer PRIMARY KEY,
    name   VARCHAR(20),
    stock  integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
