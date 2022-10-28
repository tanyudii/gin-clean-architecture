CREATE TABLE `role_permissions` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `role_serial` varchar(36) NOT NULL,
    `permission_code` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE UNIQUE INDEX uq_role_permissions_role_permission ON role_permissions (`role_serial`, `permission_code`);