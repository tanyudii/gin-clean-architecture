CREATE TABLE `user_details` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_serial` varchar(36) NOT NULL UNIQUE,
    `company_serial` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE UNIQUE INDEX uq_user_details_user_company ON user_details (`user_serial`, `company_serial`);