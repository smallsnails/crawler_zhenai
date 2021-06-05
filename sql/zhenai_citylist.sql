/*
Navicat MySQL Data Transfer

Source Server         : 1
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2021-06-05 16:49:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for zhenai_citylist
-- ----------------------------
DROP TABLE IF EXISTS `zhenai_citylist`;
CREATE TABLE `zhenai_citylist` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
