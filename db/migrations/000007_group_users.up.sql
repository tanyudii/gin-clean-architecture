CREATE TABLE `group_users` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `group_serial` varchar(36) NOT NULL,
    `user_serial` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE UNIQUE INDEX uq_group_users_group_user ON group_users (`group_serial`, `user_serial`);