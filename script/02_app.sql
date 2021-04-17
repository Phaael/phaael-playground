use appdb;

CREATE TABLE `appdb`.`accounts` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `document_number` BIGINT(20) NOT NULL,
  `available_credit_limit` DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `account_id` (`id` ASC) VISIBLE);


CREATE TABLE `appdb`.`operations_types` (
  `id` INT NOT NULL ,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));

INSERT INTO `appdb`.`operations_types` (`id`, `description`) VALUES ('1', 'COMPRA A VISTA');
INSERT INTO `appdb`.`operations_types` (`id`, `description`) VALUES ('2', 'COMPRA PARCELADA');
INSERT INTO `appdb`.`operations_types` (`id`, `description`) VALUES ('3', 'SAQUE');
INSERT INTO `appdb`.`operations_types` (`id`, `description`) VALUES ('4', 'PAGAMENTO');


CREATE TABLE `appdb`.`transactions` (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` BIGINT(20) NOT NULL,
  `operation_type_id` BIGINT(10) NOT NULL,
  `amount` DECIMAL(10,2) NOT NULL,
  `event_date` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `id_tx` (`id` ASC, `account_id` ASC) VISIBLE);


