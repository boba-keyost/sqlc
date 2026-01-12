CREATE TABLE sort_items (
                      id   BIGINT PRIMARY KEY AUTO_INCREMENT,
                      name varchar(255)      NOT NULL,
                      created DATETIME NULL
) ENGINE=InnoDB;

INSERT INTO `sort_items`(`name`, `created`) VALUES('a', '2025-01-01 20:00:00'),
('b', '2025-03-01 20:00:00'),
('c', '2025-02-01 20:00:00'),
('d', '2020-02-01 20:00:00');