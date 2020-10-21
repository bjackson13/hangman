-- MySQL Script generated by MySQL Workbench
-- Tue 20 Oct 2020 07:55:11 PM EDT
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema btj9560
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema btj9560
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `btj9560` ;
USE `btj9560` ;

-- -----------------------------------------------------
-- Table `btj9560`.`Words`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`Words` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`Words` (
  `WordId` INT NOT NULL AUTO_INCREMENT,
  `Word` CHAR(16) NULL,
  `CorrectGuesses` VARCHAR(16) NULL,
  `IncorrectGuesses` VARCHAR(25) NULL,
  PRIMARY KEY (`WordId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`User`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`User` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`User` (
  `UserId` INT NOT NULL AUTO_INCREMENT,
  `Username` VARCHAR(32) NOT NULL,
  `Password` VARCHAR(64) NOT NULL,
  `IP` VARCHAR(15) NOT NULL,
  `UserAgent` VARCHAR(45) NOT NULL,
  `LastLogin` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`UserId`),
  UNIQUE INDEX `Username_UNIQUE` (`Username` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`Chat`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`Chat` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`Chat` (
  `ChatId` INT NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`ChatId`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`Games`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`Games` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`Games` (
  `GameId` INT NOT NULL AUTO_INCREMENT,
  `WordId` INT NULL,
  `GuessingUserId` INT NULL,
  `WordCreatorId` INT NULL,
  `ChatId` INT NULL,
  PRIMARY KEY (`GameId`),
  INDEX `fk_Games_1_idx` (`WordId` ASC) VISIBLE,
  INDEX `fk_Games_guessser_idx` (`GuessingUserId` ASC) VISIBLE,
  INDEX `fk_Games_creator_idx` (`WordCreatorId` ASC) VISIBLE,
  INDEX `fk_Games_chat_idx` (`ChatId` ASC) VISIBLE,
  CONSTRAINT `fk_Games_1`
    FOREIGN KEY (`WordId`)
    REFERENCES `btj9560`.`Words` (`WordId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Games_guessser`
    FOREIGN KEY (`GuessingUserId`)
    REFERENCES `btj9560`.`User` (`UserId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Games_creator`
    FOREIGN KEY (`WordCreatorId`)
    REFERENCES `btj9560`.`User` (`UserId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Games_chat`
    FOREIGN KEY (`ChatId`)
    REFERENCES `btj9560`.`Chat` (`ChatId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`ChatUsers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`ChatUsers` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`ChatUsers` (
  `UserId` INT NOT NULL,
  `ChatId` INT NOT NULL,
  PRIMARY KEY (`UserId`, `ChatId`),
  INDEX `fk_ChatUsers_chat_idx` (`ChatId` ASC) VISIBLE,
  CONSTRAINT `fk_ChatUsers_chat`
    FOREIGN KEY (`ChatId`)
    REFERENCES `btj9560`.`Chat` (`ChatId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_ChatUsers_user`
    FOREIGN KEY (`UserId`)
    REFERENCES `btj9560`.`User` (`UserId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`Messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`Messages` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`Messages` (
  `MessageId` INT NOT NULL AUTO_INCREMENT,
  `ChatId` INT NOT NULL,
  `Timestamp` INT NULL,
  `SenderId` INT NULL,
  `MessageText` VARCHAR(256) NULL,
  PRIMARY KEY (`MessageId`, `ChatId`),
  CONSTRAINT `fk_Messages_chat`
    FOREIGN KEY (`ChatId`)
    REFERENCES `btj9560`.`Chat` (`ChatId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`Lobby`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`Lobby` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`Lobby` (
  `LobbyId` INT NOT NULL AUTO_INCREMENT,
  `ChatId` INT NOT NULL,
  PRIMARY KEY (`LobbyId`, `ChatId`),
  INDEX `fk_Lobby_chat_idx` (`ChatId` ASC) VISIBLE,
  CONSTRAINT `fk_Lobby_chat`
    FOREIGN KEY (`ChatId`)
    REFERENCES `btj9560`.`Chat` (`ChatId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `btj9560`.`LobbyUsers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `btj9560`.`LobbyUsers` ;

CREATE TABLE IF NOT EXISTS `btj9560`.`LobbyUsers` (
  `LobbyId` INT NOT NULL,
  `UserId` INT NOT NULL,
  PRIMARY KEY (`LobbyId`, `UserId`),
  INDEX `fk_LobbyUsers_user_idx` (`UserId` ASC) VISIBLE,
  CONSTRAINT `fk_LobbyUsers_user`
    FOREIGN KEY (`UserId`)
    REFERENCES `btj9560`.`User` (`UserId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_LobbyUsers_lobby`
    FOREIGN KEY (`LobbyId`)
    REFERENCES `btj9560`.`Lobby` (`LobbyId`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
