CREATE TABLE `carts` (
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `quantity` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `icon` json DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `cities` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `food_likes` (
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `food_ratings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `food_id` int NOT NULL,
  `point` float DEFAULT '0',
  `comment` text,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `food_id` (`food_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `foods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `restaurant_id` int NOT NULL,
  `category_id` int DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` float NOT NULL,
  `images` json NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `restaurant_id` (`restaurant_id`) USING BTREE,
  KEY `category_id` (`category_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `file_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `width` int NOT NULL,
  `height` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `order_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `food_origin` json DEFAULT NULL,
  `price` float NOT NULL,
  `quantity` int NOT NULL,
  `discount` float DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `order_trackings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `state` enum('waiting_for_shipper','preparing','on_the_way','delivered','cancel') NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `total_price` float NOT NULL,
  `shipper_id` int DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `shipper_id` (`shipper_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `restaurant_foods` (
  `restaurant_id` int NOT NULL,
  `food_id` int NOT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`food_id`),
  KEY `food_id` (`food_id`)
) ENGINE=InnoDB;

CREATE TABLE `restaurant_likes` (
  `restaurant_id` int NOT NULL,
  `user_id` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`restaurant_id`,`user_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB;

CREATE TABLE `restaurant_ratings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `restaurant_id` int NOT NULL,
  `point` float NOT NULL DEFAULT '0',
  `comment` text,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `restaurant_id` (`restaurant_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `restaurants` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `address` varchar(255) NOT NULL,
  `city_id` int DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json NULL,
  `logo` json NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `liked_count` int DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `user_addresses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `city_id` int NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `icon` json DEFAULT NULL,
  `address` varchar(255) NOT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `user_device_tokens` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned DEFAULT NULL,
  `is_production` tinyint(1) DEFAULT '0',
  `os` enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
  `token` varchar(255) DEFAULT NULL,
  `status` smallint unsigned NOT NULL DEFAULT '1',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `os` (`os`) USING BTREE
) ENGINE=InnoDB;

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `fb_id` varchar(50) DEFAULT NULL,
  `gg_id` varchar(50) DEFAULT NULL,
  `password` varchar(50) NOT NULL,
  `salt` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `role` enum('user','admin','shipper') NOT NULL DEFAULT 'user',
  `avatar` json DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB;

INSERT INTO `cities` (`id`, `title`, `status`, `created_at`, `updated_at`) VALUES
(1, 'An Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(2, 'V??ng T??u', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(3, 'B???c Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(4, 'B???c C???n', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(5, 'B???c Li??u', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(6, 'B???c Ninh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(7, 'B???n Tre', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(8, 'B??nh ?????nh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(9, 'B??nh D????ng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(10, 'B??nh Ph?????c', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(11, 'B??nh Thu???n', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(12, 'C?? Mau', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(13, 'C???n Th??', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(14, 'Cao B???ng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(15, '???? N???ng', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(16, '?????k L???k', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(17, '?????k N??ng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(18, '??i???n Bi??n', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(19, '?????ng Nai', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(20, '?????ng Th??p', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(21, 'Gia Lai', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(22, 'H?? Giang', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(23, 'H?? Nam', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(24, 'H?? N???i', 1, '2020-05-18 06:52:21', '2020-05-18 06:52:21'),
(25, 'H?? T??nh', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(26, 'H???i D????ng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(27, 'H???i Ph??ng', 1, '2020-05-18 06:52:22', '2020-05-18 06:52:22'),
(28, 'H???u Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(29, 'Ho?? B??nh', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(30, 'H??ng Y??n', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(31, 'Kh??nh Ho??', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(32, 'Ki??n Giang', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(33, 'Kon Tum', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(34, 'Lai Ch??u', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(35, 'L??m ?????ng', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(36, 'L???ng S??n', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(37, 'L??o Cai', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(38, 'Long An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(39, 'Nam ?????nh', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(40, 'Ngh??? An', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(41, 'Ninh B??nh', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(42, 'Ninh Thu???n', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(43, 'Ph?? Th???', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(44, 'Ph?? Y??n', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(45, 'Qu???ng B??nh', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(46, 'Qu???ng Namm', 1, '2020-05-18 06:52:23', '2020-05-18 06:52:23'),
(47, 'Qu??ng Ng??i', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(48, 'Qu??ng Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(49, 'Qu??ng Tr???', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(50, 'S??c Tr??ng', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(51, 'S??n La', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(52, 'T??y Ninh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(53, 'Th??i B??nh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(54, 'Th??i Nguy??n', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(55, 'Thanh Ho??', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(56, 'Hu???', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(57, 'Ti???n Giang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(58, 'H??? Ch?? Minh', 1, '2020-05-18 06:41:51', '2020-05-18 06:41:51'),
(59, 'Tr?? Vinh', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(60, 'Tuy??n Quang', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(61, 'V??nh Long', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(62, 'V??nh Ph??c', 1, '2020-05-18 06:55:18', '2020-05-18 06:55:18'),
(63, 'Y??n B??i', 1, '2020-05-18 06:55:19', '2020-05-18 06:55:19');