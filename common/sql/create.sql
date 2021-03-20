################################ 删除 ################################
drop table if exists `element_class`;
drop table if exists `good`;
drop table if exists `element`;
drop table if exists `main_element_attach_element_record`;
drop table if exists `element_select_size_record`;
drop table if exists `desk`;
drop table if exists `space`;
drop table if exists `order`;
drop table if exists `favor_record`;
drop table if exists `desk_class`; # 废弃
drop table if exists `space_class`;
drop table if exists `element_size_info_record`;


################################ 创建 ################################
# 桌类表
CREATE TABLE `space_class`
(
    `id`                 int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`               varchar(30)  NOT NULL DEFAULT '',
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 商品类表
CREATE TABLE `element_class`
(
    `id`                 int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`               varchar(30)  NOT NULL DEFAULT '',
    `class_type`         int          NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 元素表
CREATE TABLE `element`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `class_id`   bigint       NOT NULL DEFAULT 0,
    `name`       varchar(30)  NOT NULL DEFAULT '',
    `type`       int          NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`class_id`, `name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 元素尺寸记录表
CREATE TABLE `element_size_info_record`
(
    `id`                 int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `element_id`         bigint       NOT NULL DEFAULT 0,
    `good_id`            bigint       NOT NULL DEFAULT 0,
    `size`               varchar(30)  NOT NULL DEFAULT '小规格',
    `price`              double       NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`good_id`, `element_id`, `size`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 商品表
CREATE TABLE `good`
(
    `id`                int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `order_id`          bigint       NOT NULL DEFAULT 0,
    `main_element_id`   bigint       NOT NULL DEFAULT 0,
    `expense`           double       NOT NULL DEFAULT 0,
    `check_out_at`      timestamp    NOT NULL,
    `non_favor_expense` double       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- # 元素类表
-- CREATE TABLE `element_class`
-- (
--     `id`         int unsigned NOT NULL AUTO_INCREMENT,
--     `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     `name`       varchar(30) NOT NULL DEFAULT '',
--     PRIMARY KEY (`id`),
--     UNIQUE KEY (`name`)
-- ) ENGINE = InnoDB
--   DEFAULT CHARSET = utf8mb4
--   COLLATE = utf8mb4_0900_ai_ci;

# 元素尺寸选择记录表(如控制奶茶默认是小规格，温度默认是常温)
CREATE TABLE `element_select_size_record`
(
    `id`                  int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `good_id`             bigint       NOT NULL DEFAULT 0,
    `element_id`          bigint       NOT NULL DEFAULT 0,
    `select_size_info_id` bigint       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`good_id`, `element_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

# 主元素的附属元素记录表(如控制奶茶的温度默认是常温，而水果茶的温度默认是冷饮)
CREATE TABLE `main_element_attach_element_record`
(
    `id`                  int unsigned NOT NULL AUTO_INCREMENT,
    `created_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `good_id`             bigint       NOT NULL DEFAULT 0,
    `main_element_id`     bigint       NOT NULL DEFAULT 0,
    `attach_element_id`   bigint       NOT NULL DEFAULT 0,
    `select_size_info_id` bigint       NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`good_id`, `main_element_id`, `attach_element_id`)
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
    `space_id`          bigint       NOT NULL DEFAULT 0,
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
    `class_id`           bigint       NOT NULL DEFAULT 0,
    `name`               varchar(30)  NOT NULL DEFAULT '',
    `price`              double       NOT NULL DEFAULT 0,
    `billing_type`       int          NOT NULL DEFAULT 0,
    `picture_store_path` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`class_id`, `name`)
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
    `chargeable_object_name` varchar(30)           DEFAULT NULL,
    `chargeable_object_id`   bigint       NOT NULL DEFAULT 0,
    `favor_type`             int                   DEFAULT NULL,
    `favor_parameters`       varchar(30)  NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


################################ 数据 ################################
# 创建商品类
insert into element_class(name, picture_store_path, class_type)
values ('奶茶类', 'static/upload/温热1.jpeg', 0);
insert into element_class(name, picture_store_path, class_type)
values ('水果茶类', 'static/upload/温热1.jpeg', 0);
insert into element_class(name, picture_store_path, class_type)
values ('小吃类', 'static/upload/温热1.jpeg', 0);
insert into element_class(name, picture_store_path, class_type)
values ('饮料类', 'static/upload/温热1.jpeg', 0);
insert into element_class(name, picture_store_path, class_type)
values ('附属选项类', 'static/upload/温热1.jpeg', 0);
insert into element_class(name, picture_store_path, class_type)
values ('附属商品类', 'static/upload/温热1.jpeg', 0);

# 创建商品, 选项记录, 附属记录
insert into element(name, type, class_id)
values ('波霸奶茶', 0, 1);
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 1, '小杯', 10, 'static/upload/奶茶1.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 1, '中杯', 15, 'static/upload/奶茶2.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 1, '大杯', 15, 'static/upload/奶茶3.jpeg');
insert into element_select_size_record(good_id, element_id, select_size_info_id)
values (0, 1, 2);
insert into main_element_attach_element_record(good_id, main_element_id, attach_element_id, select_size_info_id)
values (0, 1, 2, 3);
insert into main_element_attach_element_record(good_id, main_element_id, attach_element_id, select_size_info_id)
values (0, 1, 3, 8);

# 创建附属选项元素, 选项记录
insert into element(name, type, class_id)
values ('温度', 1, 5);
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 2, '冷饮', 0, 'static/upload/温热1.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 2, '常温', 0, 'static/upload/温热1.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 2, '热饮', 0, 'static/upload/温热1.jpeg');
insert into element_select_size_record(good_id, element_id, select_size_info_id)
values (0, 2, 4);

# 创建附属商品元素, 选项记录
insert into element(name, type, class_id)
values ('珍珠', 2, 6);
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 3, '少量', 1, 'static/upload/珍珠1.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 3, '中量', 2, 'static/upload/珍珠2.jpeg');
insert into element_size_info_record(good_id, element_id, size, price, picture_store_path)
values (0, 3, '大量', 2.5, 'static/upload/珍珠3.jpeg');
insert into element_select_size_record(good_id, element_id, select_size_info_id)
values (0, 3, 7);
