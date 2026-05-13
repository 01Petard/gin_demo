/*
 Navicat Premium Dump SQL

 Source Server         : 华为云 MySQL 8.4.3
 Source Server Type    : MySQL
 Source Server Version : 80403 (8.4.3)
 Source Host           : 1.94.147.176:3306
 Source Schema         : gin_demo

 Target Server Type    : MySQL
 Target Server Version : 80403 (8.4.3)
 File Encoding         : 65001

 Date: 08/09/2025 11:41:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `phone`, `password`) VALUES (5, '2025-03-15 18:54:46.453', '2025-03-15 18:54:46.453', NULL, 'STwKpCtNSn', '13056937890', '$2a$10$4nNCx8aFk24saFJ6.el9WeSLNViu.gGbV68FNpVv.maC6ypReRkrG');
INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `phone`, `password`) VALUES (6, '2025-03-15 20:05:37.622', '2025-03-15 20:05:37.622', NULL, 'dQcykmzJDM', '13248684099', '$2a$10$B7WDt39cfN3zxH.CkQ54HuGY7drHcj.xZtcaiP4GSYLIFrvNmRsVK');
COMMIT;

-- ----------------------------
-- Table structure for users_copy1
-- ----------------------------
DROP TABLE IF EXISTS `users_copy1`;
CREATE TABLE `users_copy1` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users_copy1
-- ----------------------------
BEGIN;
INSERT INTO `users_copy1` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `phone`, `password`) VALUES (5, '2025-03-15 18:54:46.453', '2025-03-15 18:54:46.453', NULL, 'STwKpCtNSn', '13056937890', '$2a$10$4nNCx8aFk24saFJ6.el9WeSLNViu.gGbV68FNpVv.maC6ypReRkrG');
INSERT INTO `users_copy1` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `phone`, `password`) VALUES (6, '2025-03-15 20:05:37.622', '2025-03-15 20:05:37.622', NULL, 'dQcykmzJDM', '13248684099', '$2a$10$B7WDt39cfN3zxH.CkQ54HuGY7drHcj.xZtcaiP4GSYLIFrvNmRsVK');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
