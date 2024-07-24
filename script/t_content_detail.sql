CREATE TABLE `t_content_details` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `title` varchar(225) DEFAULT '' COMMENT '內容標題',
  `description` text COMMENT '內容描述',
  `author` varchar(225) DEFAULT '' COMMENT '作者',
  `video_url` varchar(225) DEFAULT '' COMMENT '視頻撥放URL',
  `thumbnail` varchar(225) DEFAULT '' COMMENT '封面圖URL',
  `catgory` varchar(225) DEFAULT '' COMMENT '內容分類',
  `duration` bigint DEFAULT 0 COMMENT '內容時長',
  `resolution` varchar(20) DEFAULT '' COMMENT '分辨率',
  `file_size` bigint DEFAULT 0 COMMENT '文件大小',
  `format` varchar(20) DEFAULT '' COMMENT '文件格式',
  `quality` int DEFAULT 0 COMMENT '影片品質',
  `approval_status` int DEFAULT 0 COMMENT '審核狀態',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
  PRIMARY KEY (`id`)
);