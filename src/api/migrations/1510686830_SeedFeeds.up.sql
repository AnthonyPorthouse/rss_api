BEGIN;

INSERT INTO `feeds` (`id`, `url`, `created_at`)
VALUES
("a7b07ca4-d8a3-40ad-8854-d37388e8d7d4", "https://feeds.feedburner.com/LICD?format=xml", NOW()),
("30f50d65-66ed-49a0-9822-1e4d77a7c462", "http://www.girlgeniusonline.com/ggmain.rss", NOW());

COMMIT;