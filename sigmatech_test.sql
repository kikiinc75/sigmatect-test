-- MySQL dump 10.13  Distrib 8.0.30, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: sigmatect_db
-- ------------------------------------------------------
-- Server version	5.7.39
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
/*!50503 SET NAMES utf8 */
;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;
--
-- Table structure for table `transactions`
--
DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `contract_number` varchar(225) DEFAULT NULL,
  `otr` double DEFAULT NULL,
  `admin_fee` double DEFAULT NULL,
  `installment` int(11) DEFAULT NULL,
  `interest_amount` double DEFAULT NULL,
  `asset_name` varchar(225) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `transactions`
--
LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */
;
INSERT INTO `transactions`
VALUES (
    1,
    1,
    'ct001',
    5000,
    1000,
    1,
    5000,
    'baju',
    '2023-02-07 08:19:08',
    '2023-02-07 08:19:08'
  );
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `user_limits`
--
DROP TABLE IF EXISTS `user_limits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `user_limits` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `tenor` int(11) NOT NULL,
  `amount` double DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 9 DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `user_limits`
--
LOCK TABLES `user_limits` WRITE;
/*!40000 ALTER TABLE `user_limits` DISABLE KEYS */
;
INSERT INTO `user_limits`
VALUES (1, 1, 1, 89000, NULL, '2023-02-07 08:19:08'),
(2, 1, 2, 200000, NULL, NULL),
(3, 1, 3, 500000, NULL, NULL),
(4, 1, 6, 700000, NULL, NULL),
(5, 2, 1, 1000000, NULL, NULL),
(6, 2, 2, 1200000, NULL, NULL),
(7, 2, 3, 1500000, NULL, NULL),
(8, 2, 6, 2000000, NULL, NULL);
/*!40000 ALTER TABLE `user_limits` ENABLE KEYS */
;
UNLOCK TABLES;
--
-- Table structure for table `users`
--
DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
/*!50503 SET character_set_client = utf8mb4 */
;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nik` varchar(255) DEFAULT NULL,
  `legal_name` varchar(255) DEFAULT NULL,
  `birth_place` varchar(255) DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `salary` double DEFAULT NULL,
  `photos_card` varchar(255) DEFAULT NULL,
  `photos_selfie` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 14 DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */
;
--
-- Dumping data for table `users`
--
LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */
;
INSERT INTO `users`
VALUES (
    1,
    'Budi',
    'budi@example.com',
    '$2a$04$kwistoKeorYgifu3gJJnzu4KS7XwW5ywRdiYObRY9fHubjycSpj4W',
    '1234567890123456',
    'Budi',
    'Ponorogo',
    '2020-02-22',
    20000,
    'test',
    'test',
    NULL,
    '2023-02-07 08:54:23'
  ),
(
    2,
    'Annisa',
    'annisa@example.com',
    '$2a$04$kwistoKeorYgifu3gJJnzu4KS7XwW5ywRdiYObRY9fHubjycSpj4W',
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL
  ),
(
    3,
    'iqbal',
    'iqbal@gmail.com',
    '$2a$04$kwistoKeorYgifu3gJJnzu4KS7XwW5ywRdiYObRY9fHubjycSpj4W',
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    '2023-02-07 07:09:20',
    '2023-02-07 07:09:20'
  );
/*!40000 ALTER TABLE `users` ENABLE KEYS */
;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;
-- Dump completed on 2023-02-07 23:41:00