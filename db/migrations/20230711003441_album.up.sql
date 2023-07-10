-- +migrate Up
CREATE TABLE `albums` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NULL,
    `contents` TEXT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,
    `deleted_at` DATETIME(3) NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;