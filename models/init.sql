DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `username` VARCHAR(255) DEFAULT '',
  `password` VARCHAR(255) DEFAULT '',
  `fullname`  VARCHAR(255) DEFAULT '',
  `phone_number` VARCHAR(255) DEFAULT '',
  `number_video` INT(10)  DEFAULT 0,
  `total_size` INT(10)  DEFAULT 0,
  `max_size` INT(10)  DEFAULT 2100,
  `created_at`  DATETIME    DEFAULT NOW(),
  `deleted_at`                 DATETIME     DEFAULT NULL

);

DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `id_user` BIGINT(20)  NOT NULL,
  `mac` VARCHAR(255) NOT NULL,
  `video_name` VARCHAR(255) DEFAULT '',
  `video_size` BIGINT(20)  NOT NULL,
  `video_time` BIGINT(20)  NOT NULL,
  `status` TINYINT(3) DEFAULT 0,
  `expired` DATETIME     DEFAULT NULL,
  `location` VARCHAR(255) DEFAULT '',
  `created_at`                 DATETIME    DEFAULT NOW(),
  `updated_at`                 DATETIME    DEFAULT NOW() ON UPDATE NOW(),
  `deleted_at`                 DATETIME    DEFAULT NULL
);

DROP TABLE IF EXISTS `medias`;
CREATE TABLE `medias` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `id_user` BIGINT(20)  NOT NULL,
  `video_name` VARCHAR(255) DEFAULT '',
  `video_size` BIGINT(20)  NOT NULL,
  `video_time` BIGINT(20)  NOT NULL,
  `created_at`                 DATETIME    DEFAULT NOW(),
  `deleted_at`                 DATETIME    DEFAULT NULL
);

DROP TABLE IF EXISTS `authenkey`;
CREATE TABLE `authenkey` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `key` VARCHAR(255) NOT NULL,
  `created_at`                 DATETIME    DEFAULT NOW()
);
