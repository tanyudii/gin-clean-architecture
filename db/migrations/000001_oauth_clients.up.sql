CREATE TABLE `oauth_clients` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `name` varchar(255) NOT NULL,
     `is_active` tinyint(1) DEFAULT 1,
     `is_internal` tinyint(1) DEFAULT 0,
     `client_id` varchar(255) NOT NULL UNIQUE,
     `client_secret` varchar(255) NOT NULL,
     `domain` varchar(255) DEFAULT NULL,
     `user_id` VARCHAR(255) DEFAULT NULL,
     `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
     `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
     `created_by` varchar(36) DEFAULT NULL,
     `updated_by` varchar(36) DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
