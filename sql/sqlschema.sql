-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: btj9560
-- ------------------------------------------------------
-- Server version	8.0.22-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Games`
--

DROP TABLE IF EXISTS `Games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Games` (
  `GameId` int NOT NULL AUTO_INCREMENT,
  `WordId` int DEFAULT NULL,
  `GuessingUserId` int NOT NULL,
  `WordCreatorId` int NOT NULL,
  `PendingGuess` varchar(1) DEFAULT NULL,
  PRIMARY KEY (`GameId`),
  UNIQUE KEY `WordCreatorId` (`WordCreatorId`),
  UNIQUE KEY `GuessingUserId` (`GuessingUserId`),
  KEY `fk_Games_1_idx` (`WordId`),
  KEY `fk_Games_guessser_idx` (`GuessingUserId`),
  KEY `fk_Games_creator_idx` (`WordCreatorId`),
  CONSTRAINT `fk_Games_1` FOREIGN KEY (`WordId`) REFERENCES `Words` (`WordId`),
  CONSTRAINT `fk_Games_creator` FOREIGN KEY (`WordCreatorId`) REFERENCES `User` (`UserId`),
  CONSTRAINT `fk_Games_guessser` FOREIGN KEY (`GuessingUserId`) REFERENCES `User` (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Games`
--

LOCK TABLES `Games` WRITE;
/*!40000 ALTER TABLE `Games` DISABLE KEYS */;
/*!40000 ALTER TABLE `Games` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Lobby`
--

DROP TABLE IF EXISTS `Lobby`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Lobby` (
  `LobbyId` int NOT NULL,
  `ChatId` int NOT NULL,
  PRIMARY KEY (`LobbyId`,`ChatId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Lobby`
--

LOCK TABLES `Lobby` WRITE;
/*!40000 ALTER TABLE `Lobby` DISABLE KEYS */;
INSERT INTO `Lobby` VALUES (1,-404);
/*!40000 ALTER TABLE `Lobby` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `LobbyUsers`
--

DROP TABLE IF EXISTS `LobbyUsers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `LobbyUsers` (
  `UserId` int NOT NULL,
  `PendingInviteId` int DEFAULT NULL,
  PRIMARY KEY (`UserId`),
  KEY `fk_LobbyUsers_user_idx` (`UserId`),
  KEY `fk_LobbyUsers_invite_idx` (`PendingInviteId`),
  CONSTRAINT `fk_LobbyUsers_invite` FOREIGN KEY (`PendingInviteId`) REFERENCES `User` (`UserId`),
  CONSTRAINT `fk_LobbyUsers_user` FOREIGN KEY (`UserId`) REFERENCES `User` (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `LobbyUsers`
--

LOCK TABLES `LobbyUsers` WRITE;
/*!40000 ALTER TABLE `LobbyUsers` DISABLE KEYS */;
/*!40000 ALTER TABLE `LobbyUsers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Messages`
--

DROP TABLE IF EXISTS `Messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Messages` (
  `MessageId` int NOT NULL AUTO_INCREMENT,
  `ChatId` int NOT NULL,
  `Timestamp` int DEFAULT NULL,
  `SenderId` int DEFAULT NULL,
  `MessageText` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`MessageId`,`ChatId`),
  KEY `fk_Messages_sender_idx` (`SenderId`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Messages`
--

LOCK TABLES `Messages` WRITE;
/*!40000 ALTER TABLE `Messages` DISABLE KEYS */;
INSERT INTO `Messages` VALUES (12,-404,1604352775,2,'This is the first message'),(13,-404,1604352777,3,'I am the postiest man'),(14,-404,1604352780,8,'Hello world'),(15,-404,1604352785,8,'Here is another hello'),(16,-404,1604352790,9,'JIIIMMMYYYY'),(17,-404,1604352795,10,'<script>alert(\'I am an asshole\');</script>');
/*!40000 ALTER TABLE `Messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `User` (
  `UserId` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(32) NOT NULL,
  `Password` varchar(64) NOT NULL,
  `IP` varchar(15) NOT NULL,
  `UserAgent` varchar(256) NOT NULL,
  `LastLogin` int unsigned NOT NULL,
  PRIMARY KEY (`UserId`),
  UNIQUE KEY `Username_UNIQUE` (`Username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (2,'bren','$2a$10$fOJNV39k723Rx0czsUA0ZuDTXE5oL59x83u4sdjL..oo8q6VCFj.W','127.0.0.1','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:82.0) Gecko/20100101 Firefox/82.0',1604350992),(3,'postman','$2a$10$dd//wUTHGVWGlpUF5NuNvuliuuGjPDu6NgJGI5ysI5..nTgiC/me.','127.0.0.1','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:82.0) Gecko/20100101 Firefox/82.0',1604352264),(8,'timmy','$2a$10$ZctnN9f.DD383miy6dbAsuUfhETjiA./ZqC6UEEymj1gSo41ycn4W','::1','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36',1604351122),(9,'user','$2a$10$VDv2TwdtYqqX.dlXFNYQA.rngEU2ibpFPJLmfypHJ.TGMYhz5UFBu','127.0.0.1','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:82.0) Gecko/20100101 Firefox/82.0',1604351783),(10,'jimmy','$2a$10$8iQ2w6pVu7jkKxx9qa9V9uxY/vvPEIHNTJeIs7oPpLUwvhcn9kH5S','::1','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36',1604351339);
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Words`
--

DROP TABLE IF EXISTS `Words`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Words` (
  `WordId` int NOT NULL AUTO_INCREMENT,
  `WordLength` int DEFAULT NULL,
  `CorrectGuesses` varchar(32) DEFAULT NULL,
  `IncorrectGuesses` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`WordId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Words`
--

LOCK TABLES `Words` WRITE;
/*!40000 ALTER TABLE `Words` DISABLE KEYS */;
INSERT INTO `Words` VALUES (1,2,NULL,NULL);
/*!40000 ALTER TABLE `Words` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-02 16:55:53
