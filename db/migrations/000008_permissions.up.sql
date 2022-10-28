CREATE TABLE `permissions` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `code` varchar(255) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `transaction` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE UNIQUE INDEX uq_permissions_transaction_name ON permissions (`transaction`, `name`);