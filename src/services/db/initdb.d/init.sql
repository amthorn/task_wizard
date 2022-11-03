CREATE DATABASE IF NOT EXISTS task_wizard;
USE task_wizard;

CREATE TABLE `projects` (
    `id` BIGINT auto_increment PRIMARY KEY,
    `name` VARCHAR(255)
);

INSERT INTO `projects` (`id`, `name`) VALUES (1, "First"), (2, "Second"), (3, "Third");