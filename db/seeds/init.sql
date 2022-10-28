INSERT INTO `oauth_clients` (`id`, `serial`, `name`, `is_internal`, `client_id`, `client_secret`)
VALUES
    (1, "b610a230-9956-49b6-a2e6-21ca6d05b8bc", "Vodea Cloud", 1, "fc93c408-877a-48f3-af66-84ee2ce6a5c5", "$2a$10$wbgvM2GbFW3CcFj1UQTf8ewyogNCQF0/nPSXLeqp3n8JhljYwlUUC"), -- Secret: itsverysecret
    (2, "92faf0f8-1315-4877-bf5c-e069dc324848", "External", 0, "737258ac-81f0-4c06-9c18-95b86aeb8b81", "$2a$10$UyPM.sGS8B3IunQkSBnzGunj2vqJ3f9u1rpkAUayfDc0CxxzjIfC2") -- Secret: itsverysecret
ON DUPLICATE KEY UPDATE `name` = VALUES(`name`), `client_id` = VALUES(`client_id`), `client_secret` = VALUES(`client_secret`);