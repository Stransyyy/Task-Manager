 show databases;

CREATE DATABASE IF NOT EXISTS tasks;

use tasks;

show databases;

CREATE TABLE IF NOT EXISTS `storage`(
    taskID int auto_increment primary key,
    title varchar (100) NOT NULL,
    `description` varchar(250),
    due_date time,
    completed bool);

show tables;

describe `storage`;

Alter table `storage`
    MODIFY COLUMN due_date time NOT NULL,
    MODIFY COLUMN completed bool NOT NULL;