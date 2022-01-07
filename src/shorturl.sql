/*
 Navicat MySQL Data Transfer

 Source Server         : 阿里云RDS -2022 11-1到期
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : 8.142.199.134:3306
 Source Schema         : shorturl

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 07/01/2022 19:09:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for urlinfos
-- ----------------------------
DROP TABLE IF EXISTS `urlinfos`;
CREATE TABLE `urlinfos`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `surl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `lurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `views` int(11) NOT NULL,
  `createTime` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of urlinfos
-- ----------------------------
INSERT INTO `urlinfos` VALUES (3, 'https://www.xingly.cn', '9Ep8u2JLPH1', 1, '2022-01-07 10:56:57');
INSERT INTO `urlinfos` VALUES (4, 'https://ww.xingly.cn', 'BhrHvG6qVaa', 0, '2022-01-07 11:02:52');
INSERT INTO `urlinfos` VALUES (5, 'https:/\\/ww.xingly.cn', 'BFUnJgCdoy', 0, '2022-01-07 11:03:05');

SET FOREIGN_KEY_CHECKS = 1;
