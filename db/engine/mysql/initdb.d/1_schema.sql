DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
use app;

-- ---- drop ----
-- DROP TABLE IF EXISTS `bans`;
-- ---- create ----
-- create table IF not exists `bans`
-- (
--   `id`               INT(20) AUTO_INCREMENT,
--   `name`             VARCHAR(20) NOT NULL,
--   `stock`            int,
--   PRIMARY KEY (`id`)
-- ) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


---- drop ----
DROP TABLE IF EXISTS `patties`;

---- create ----
CREATE TABLE `patties` (
  `id`         int(20) NOT NULL AUTO_INCREMENT,
  `name`       varchar(20) COLLATE utf8_bin NOT NULL,
  `stock`      int(5) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ---- drop ----
-- DROP TABLE IF EXISTS `vegetables`;

-- ---- create ----
-- create table IF not exists `vegetables`
-- (
--   `id`               INT(20) AUTO_INCREMENT,
--   `name`             VARCHAR(20) NOT NULL,
--   `stock`            int,
--   PRIMARY KEY (`id`)
-- ) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


-- ---- drop ----
-- DROP TABLE IF EXISTS `ingredients`;

-- ---- create ----
-- create table IF not exists `ingredients`
-- (
--   `id`               INT(20) AUTO_INCREMENT,
--   `name`             VARCHAR(20) NOT NULL,
--   `stock`            int,
--   PRIMARY KEY (`id`)
-- ) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

