-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         10.5.7-MariaDB - mariadb.org binary distribution
-- SO del servidor:              Win64
-- HeidiSQL Versión:             11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Volcando estructura de base de datos para controlhorarios
CREATE DATABASE IF NOT EXISTS `controlhorarios` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `controlhorarios`;

-- Volcando estructura para tabla controlhorarios.company
CREATE TABLE IF NOT EXISTS `company` (
  `company_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `company_name` varchar(50) NOT NULL DEFAULT 'Company without name',
  `created_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `location` varchar(280) DEFAULT NULL,
  PRIMARY KEY (`company_id`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `FK_company_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.company: ~3 rows (aproximadamente)
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
INSERT INTO `company` (`company_id`, `user_id`, `company_name`, `created_date`, `location`) VALUES
	(1, 4, 'Tailandia Records SL', '2020-11-13 12:24:03', NULL),
	(2, 5, 'Macrohard ', '2020-11-13 12:37:28', NULL),
	(3, 7, 'Ubihard', '2020-11-13 12:42:50', NULL),
	(4, 8, 'Riote Games', '2020-11-16 13:20:49', NULL);
/*!40000 ALTER TABLE `company` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.employee
CREATE TABLE IF NOT EXISTS `employee` (
  `employee_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `first_name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT NULL,
  `profile_picture` blob DEFAULT NULL,
  `project_id` int(11) unsigned DEFAULT NULL,
  `company_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`employee_id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `FK2_employee_projects` (`project_id`),
  KEY `FK1_user_id` (`user_id`) USING BTREE,
  KEY `FK3_employee_company` (`company_id`),
  CONSTRAINT `FK1_employee_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `FK2_employee_projects` FOREIGN KEY (`project_id`) REFERENCES `projects` (`project_id`),
  CONSTRAINT `FK3_employee_company` FOREIGN KEY (`company_id`) REFERENCES `company` (`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.employee: ~3 rows (aproximadamente)
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
INSERT INTO `employee` (`employee_id`, `user_id`, `first_name`, `last_name`, `profile_picture`, `project_id`, `company_id`) VALUES
	(1, 4, 'airan', 'a', NULL, NULL, 1),
	(3, 6, 'Jaime', 'Jaime', NULL, NULL, 2),
	(9, 9, 'Yodra', 'Lopez', NULL, NULL, 3);
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.employee_records
CREATE TABLE IF NOT EXISTS `employee_records` (
  `record_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(280) DEFAULT NULL,
  `start_time` timestamp NULL DEFAULT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `employee_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`record_id`),
  KEY `FK1_record_employee` (`employee_id`) USING BTREE,
  CONSTRAINT `FK1_records_employee` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`employee_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.employee_records: ~2 rows (aproximadamente)
/*!40000 ALTER TABLE `employee_records` DISABLE KEYS */;
INSERT INTO `employee_records` (`record_id`, `description`, `start_time`, `end_time`, `employee_id`) VALUES
	(1, 'Trabajo', '2020-11-13 12:39:00', NULL, 1),
	(2, 'Trabajo', '2020-11-13 12:41:44', NULL, 3),
	(3, 'Trabajo', '2020-11-13 12:42:07', '2020-11-13 13:42:01', 3),
	(4, 'Dar clases', '2020-11-16 13:30:12', NULL, 9),
	(5, 'Finalizar clases', '2020-11-16 13:30:44', '2020-11-16 18:30:35', 9);
/*!40000 ALTER TABLE `employee_records` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.holidays
CREATE TABLE IF NOT EXISTS `holidays` (
  `holiday_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `holiday_title` varchar(100) NOT NULL DEFAULT 'Holiday title',
  `holiday_date` date NOT NULL DEFAULT curdate(),
  `company_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`holiday_id`),
  KEY `FK1_holiday_company` (`company_id`),
  CONSTRAINT `FK1_holiday_company` FOREIGN KEY (`company_id`) REFERENCES `company` (`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.holidays: ~2 rows (aproximadamente)
/*!40000 ALTER TABLE `holidays` DISABLE KEYS */;
INSERT INTO `holidays` (`holiday_id`, `holiday_title`, `holiday_date`, `company_id`) VALUES
	(1, 'Navidad', '2020-11-13', 1),
	(2, 'Semana santa', '2020-11-16', 3),
	(3, 'Carnavales', '2021-02-10', 4);
/*!40000 ALTER TABLE `holidays` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.projects
CREATE TABLE IF NOT EXISTS `projects` (
  `project_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_name` varchar(50) NOT NULL DEFAULT 'Project without name',
  `company_id` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`project_id`),
  KEY `FK1_proyects_company` (`company_id`),
  CONSTRAINT `FK1_proyects_company` FOREIGN KEY (`company_id`) REFERENCES `company` (`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.projects: ~1 rows (aproximadamente)
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` (`project_id`, `project_name`, `company_id`) VALUES
	(2, 'Holaluz', 2),
	(3, 'Control de horarios', 1),
	(5, 'Project scorpio', 4);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.sections
CREATE TABLE IF NOT EXISTS `sections` (
  `section_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `section_name` varchar(100) NOT NULL DEFAULT 'Department name',
  `company_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`section_id`),
  KEY `FK1_sections_company` (`company_id`),
  CONSTRAINT `FK1_sections_company` FOREIGN KEY (`company_id`) REFERENCES `company` (`company_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.sections: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `sections` DISABLE KEYS */;
INSERT INTO `sections` (`section_id`, `section_name`, `company_id`) VALUES
	(1, 'IT', 1),
	(2, 'RRHH', 1),
	(3, 'IT', 2),
	(4, 'RRHH', 2),
	(5, 'IT', 3),
	(6, 'RRHH', 3);
/*!40000 ALTER TABLE `sections` ENABLE KEYS */;

-- Volcando estructura para tabla controlhorarios.users
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '',
  `email` varchar(50) NOT NULL DEFAULT '',
  `password` varchar(50) NOT NULL DEFAULT '',
  `joined_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `rol` enum('company','employee') NOT NULL DEFAULT 'employee',
  `status` boolean DEFAULT TRUE,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

-- Volcando datos para la tabla controlhorarios.users: ~3 rows (aproximadamente)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`user_id`, `username`, `email`, `password`, `joined_date`, `rol`) VALUES
	(4, 'Airan', 'airan@gmail.com', '12345', '2020-11-13 11:49:16', 'employee'),
	(5, 'Ariane', 'ariane@gmail.com', '12345', '2020-11-13 12:37:08', 'company'),
	(6, 'Jaime', 'jaime@gmail.com', '12345', '2020-11-13 12:40:49', 'employee'),
	(7, 'Manuel', 'Manuel@gmail.com', '12345', '2020-11-13 12:42:30', 'company'),
	(8, 'Riote', 'riot@gmail.com', '12345', '2020-11-16 13:20:26', 'company'),
	(9, 'Yodra', 'yodra@gmail.com', '12345', '2020-11-16 13:29:23', 'employee');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
