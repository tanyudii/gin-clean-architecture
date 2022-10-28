CREATE TABLE `companies` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `serial` varchar(36) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `initial` varchar(100) DEFAULT NULL,
    `email` varchar(255) DEFAULT NULL,
    `phone` varchar(100) DEFAULT NULL,
    `fax` varchar(100) DEFAULT NULL,
    `address` varchar(255) DEFAULT NULL,
    `city` varchar(255) DEFAULT NULL,
    `province` varchar(255) DEFAULT NULL,
    `zip_code` varchar(5) DEFAULT NULL,
    `is_active` tinyint(1) DEFAULT 1,
    `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `created_by` varchar(36) DEFAULT NULL,
    `updated_by` varchar(36) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_companies_name ON companies (`name`);