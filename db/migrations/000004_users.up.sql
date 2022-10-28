CREATE TABLE `users` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `serial` varchar(36) NOT NULL UNIQUE,
   `name` varchar(255) NOT NULL,
   `email` varchar(255) NOT NULL UNIQUE,
   `password` varchar(255) NOT NULL,
   `phone` varchar(100) DEFAULT NULL UNIQUE,
   `avatar_url` varchar(255) DEFAULT NULL,
   `type` enum('ADMIN', 'USER') NOT NULL DEFAULT 'ADMIN',
   `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
   `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
   `created_by` varchar(36) DEFAULT NULL,
   `updated_by` varchar(36) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_users_email_phone_name ON users (`email`, `phone`, `name`);