CREATE TABLE `notifications` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `serial` varchar(36) NOT NULL UNIQUE,
    `user_serial` varchar(36) NOT NULL,
    `data` TEXT DEFAULT NULL,
    `read_at` datetime(6) DEFAULT NULL,
    `created_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_notifications_user_serial_read ON notifications (`user_serial`, `read_at`);