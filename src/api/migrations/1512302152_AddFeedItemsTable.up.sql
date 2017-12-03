BEGIN;

CREATE TABLE `items`
(
  `id` CHAR(36) NOT NULL,
  `feed_id` CHAR(36) NOT NULL,

  `link` VARCHAR(255) NOT NULL,

  `content` VARCHAR(255) NOT NULL,

  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME,

  PRIMARY KEY (`id`),
  FOREIGN KEY (`feed_id`)
    REFERENCES `feeds`(`id`)
    ON DELETE CASCADE
);

COMMIT;