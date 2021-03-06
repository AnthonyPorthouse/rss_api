START TRANSACTION;

CREATE TABLE `feeds`
(
  `id` CHAR(36) NOT NULL,
  `url` VARCHAR(255) NOT NULL,

  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME,

  PRIMARY KEY (`id`),
  UNIQUE KEY url (`url`)
);

COMMIT;