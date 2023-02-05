CREATE TABLE `douyin`.`user_logins`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `username` varchar(32) NOT NULL COMMENT '登录用户名',
  `password` varchar(128) NOT NULL COMMENT '登录密码',
  `created_at` datetime(3) NULL COMMENT 'gorm',
  `updated_at` datetime(3) NULL COMMENT 'gorm',
  `deleted_at` datetime(3) NULL COMMENT 'gorm',
  PRIMARY KEY (`id`)
) COMMENT = '用户登录信息表';
