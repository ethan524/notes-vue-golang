CREATE DATABASE `notes` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

/************************************************/
/*
    作用：用户表
    gender:1=男,2=女
    status:0=未激活状态，1=激活状态，2=禁止状态
 */
/************************************************/
CREATE TABLE user
(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL DEFAULT "" COMMENT "用户名称",
    `password` VARCHAR(128) NOT NULL DEFAULT "" COMMENT "用户密码",
    `email` VARCHAR(128) NOT NULL DEFAULT "" COMMENT "邮箱",
    `gender` CHAR(1) NOT NULL DEFAULT "1" COMMENT "性别",
    `telphone` VARCHAR(11) NOT NULL DEFAULT "" COMMENT "手机号码",
    `thumnail` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "头像图",
    `introduction` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "简介",
    `status` CHAR(1) NOT NULL DEFAULT "" COMMENT "状态",
    `create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
    `update_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT "更新时间",
    `role` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "角色ID",
    PRIMARY KEY(`id`),
    UNIQUE(`email`)
)engine=innodb default charset=utf8 comment 'user';