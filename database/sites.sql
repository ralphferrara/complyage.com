CREATE TABLE `myapp`.`sites` (
  `id_site` INT NOT NULL AUTO_INCREMENT,
  `site_name` VARCHAR(128) NULL,
  `site_url` VARCHAR(255) NULL,
  `site_status` CHAR(4) NULL,
  PRIMARY KEY (`id_site`),
  INDEX `site_url` (`site_url` ASC),
  INDEX `site_status` (`site_status` ASC));
