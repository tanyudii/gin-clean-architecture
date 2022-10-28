CREATE TABLE `role_users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `role_serial` varchar(36) NOT NULL,
    `user_serial` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE UNIQUE INDEX uq_role_users_role_user ON role_users (`role_serial`, `user_serial`);