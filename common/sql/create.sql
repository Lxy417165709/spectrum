# 元素表
CREATE TABLE `element`
(
    `id`                 int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`               varchar(255) NOT NULL DEFAULT '',
    `type`               int          NOT NULL DEFAULT 0,
    `class_name`         varchar(255) NOT NULL DEFAULT '未分类',
    `size`               varchar(255) NOT NULL DEFAULT '小规格',
    `price`              double       NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`class_name`, `name`, `size`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 商品表
CREATE TABLE `good`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`              varchar(255) NOT NULL DEFAULT '',
    `order_id`          bigint       NOT NULL DEFAULT 0,
    `expense`           double       NOT NULL DEFAULT 0,
    `check_out_at`      timestamp    NOT NULL,
    `non_favor_expense` double       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 商品类表
CREATE TABLE `good_class`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`       varchar(255) NOT NULL DEFAULT '',
    `class_type` int          NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- # 元素类表
-- CREATE TABLE `element_class`
-- (
--     `id`         int unsigned NOT NULL AUTO_INCREMENT,
--     `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     `name`       varchar(255) NOT NULL DEFAULT '',
--     PRIMARY KEY (`id`),
--     UNIQUE KEY (`name`)
-- ) ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COLLATE = utf8mb4_0900_ai_ci;

# 主元素的附属元素记录表
CREATE TABLE `main_element_attach_element_record`
(
    `id`                  int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `main_element_name`   varchar(255) NOT NULL DEFAULT '',
    `good_id`             bigint       NOT NULL DEFAULT 0,
    `attach_element_name` varchar(255) NOT NULL DEFAULT '',
    `select_size`         varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`good_id`, `main_element_name`, `attach_element_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 主元素尺寸记录表
CREATE TABLE `main_element_size_record`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `main_element_name` varchar(255) NOT NULL DEFAULT '',
    `good_id`           bigint       NOT NULL DEFAULT 0,
    `select_size`       varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`good_id`, `main_element_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 桌子表
CREATE TABLE `desk`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `start_at`          timestamp    NOT NULL,
    `end_at`            timestamp    NOT NULL,
    `session_count`     bigint       NOT NULL DEFAULT 0,
    `space_name`        varchar(255) NOT NULL DEFAULT '',
    `space_class_name`  varchar(255) NOT NULL DEFAULT '',
    `expense`           double       NOT NULL DEFAULT 0,
    `check_out_at`      timestamp    NOT NULL,
    `non_favor_expense` double       NOT NULL DEFAULT 0,
    `order_id`          bigint       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 空间表
CREATE TABLE `space`
(
    `id`                 int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`               varchar(255) NOT NULL DEFAULT '',
    `class_name`         varchar(255) NOT NULL DEFAULT '',
    `price`              double       NOT NULL DEFAULT 0,
    `billing_type`       int          NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`class_name`, `name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 订单表
CREATE TABLE `order`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `expense`           double       NOT NULL DEFAULT 0,
    `check_out_at`      timestamp    NOT NULL,
    `non_favor_expense` double       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 优惠表
CREATE TABLE `favor_record`
(
    `id`                     int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`             timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `chargeable_object_name` varchar(255)          DEFAULT NULL,
    `chargeable_object_id`   bigint       NOT NULL DEFAULT 0,
    `favor_type`             int                   DEFAULT NULL,
    `favor_parameters`       varchar(255) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci