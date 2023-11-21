SET FOREIGN_KEY_CHECKS = 0;
SELECT Concat('TRUNCATE TABLE `', table_schema, '`.`', TABLE_NAME, '`;')
INTO OUTFILE '../truncate_tables.sql'
FROM INFORMATION_SCHEMA.TABLES
WHERE table_schema = 'scanner';
SET FOREIGN_KEY_CHECKS = 1;