CREATE TABLE `notification_configs` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_serial` varchar(36) NOT NULL,
    `code` VARCHAR(255) NOT NULL,
    `enabled` tinyint(1) DEFAULT 1,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_notification_configs_user_serial ON notification_configs (`user_serial`, `code`);