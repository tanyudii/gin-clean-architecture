CREATE TABLE `groups` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `serial` varchar(36) NOT NULL UNIQUE,
    `company_serial` varchar(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `is_active` tinyint(1) DEFAULT 1,
    `parent_serial` varchar(36) DEFAULT NULL,
    `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `created_by` varchar(36) DEFAULT NULL,
    `updated_by` varchar(36) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_groups_company_name ON groups (`company_serial`, `name`);