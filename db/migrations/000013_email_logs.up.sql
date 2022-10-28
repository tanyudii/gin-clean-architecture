CREATE TABLE `email_logs` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `company_serial` varchar(36) NOT NULL,
    `to_email` varchar(255) NOT NULL,
    `to_name` varchar(255) DEFAULT NULL,
    `from_email` varchar(255) NOT NULL,
    `from_name` varchar(255) DEFAULT NULL,
    `subject` varchar(255) NOT NULL,
    `provider` varchar(255) NOT NULL,
    `template` varchar(255) NOT NULL,
    `template_code` varchar(255) NOT NULL,
    `data` TEXT DEFAULT NULL,
    `accepted_at` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `delivered_at` datetime(6) DEFAULT NULL,
    `retry` INT(8) DEFAULT 1,
    `max_retry` INT(8) DEFAULT 1,
    `last_error` TEXT DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE INDEX idx_email_logs_company_serial ON email_logs (`company_serial`)
CREATE INDEX idx_email_logs_company_to_subject ON email_logs (`company_serial`, `to_email`, `to_name`);
CREATE INDEX idx_email_logs_company_from_subject ON email_logs (`company_serial`, `from_email`, `from_name`);
CREATE INDEX idx_email_logs_company_template ON email_logs (`company_serial`, `template`);
CREATE INDEX idx_email_logs_company_template_code ON email_logs (`company_serial`, `template_code`);