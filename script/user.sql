CREATE DATABASE `cms_account`

CREATE TABLE `account` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `user_id` varchar(64) DEFAULT '' COMMENT '用戶ID',
  `password` varchar(64) DEFAULT '' COMMENT '密碼',
  `nickname` varchar(64) DEFAULT '' COMMENT '暱稱',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci