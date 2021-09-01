DROP TABLE IF exists `bookinfo`;
CREATE TABLE `bookinfo` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
                            `author` varchar(64) NOT NULL COMMENT '书籍作者',
                            `book_name` varchar(64) NOT NULL COMMENT '书籍名称',
                            `size` float(64) NOT NULL COMMENT '书籍大小MB',
                            `level` int(64) NOT NULL COMMENT '书籍评分等级',
                            `time` varchar(64) DEFAULT NULL COMMENT '书籍上传时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

