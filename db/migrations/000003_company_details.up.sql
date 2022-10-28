CREATE TABLE `company_details` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `company_serial` varchar(36) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
