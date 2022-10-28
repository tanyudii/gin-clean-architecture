CREATE TABLE `email_templates` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `code` varchar(255) NOT NULL UNIQUE,
    `provider` varchar(255) NOT NULL,
    `template` varchar(255) NOT NULL,
    `from_email` varchar(255) NOT NULL,
    `from_name` varchar(255) DEFAULT NULL,
    `default_subject` varchar(255) DEFAULT NULL,
    `default_data` TEXT DEFAULT NULL,
    `default_retry` INT(8) DEFAULT NULL,
    `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `created_by` varchar(36) DEFAULT NULL,
    `updated_by` varchar(36) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
