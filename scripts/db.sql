-- 创建gin_demo数据库
CREATE DATABASE `scanner` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建gin_demo数据库的用户
CREATE USER 'scan'@'%' IDENTIFIED WITH mysql_native_password BY 'scan1688@#' PASSWORD EXPIRE NEVER;

-- 授权gin_demo数据库的用户
GRANT ALL PRIVILEGES ON scanner.* TO 'scan'@'%';

-- 刷新权限
FLUSH PRIVILEGES;

-- 查询数据中的用户
SELECT * FROM mysql.user

-- 清空表，并重置记数器
TRUNCATE TABLE

-- 清空库中所有表
SET FOREIGN_KEY_CHECKS = 0;
SELECT Concat('TRUNCATE TABLE `', table_schema, '`.`', TABLE_NAME, '`;')
INTO OUTFILE '../truncate_tables.sql'
FROM INFORMATION_SCHEMA.TABLES
WHERE table_schema = 'scanner';
SET FOREIGN_KEY_CHECKS = 1;

-- 清表脚本
mysql -u scan -p scanner < ../truncate_tables.sql
