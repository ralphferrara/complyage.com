CREATE TABLE `users` (
  `id_user` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_email` varchar(160) DEFAULT NULL,
  `user_password` varchar(255) DEFAULT NULL,
  `user_status` varchar(4) DEFAULT NULL,
  `user_merchant` tinyint DEFAULT '0',
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `uni_users_email` (`user_password`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
