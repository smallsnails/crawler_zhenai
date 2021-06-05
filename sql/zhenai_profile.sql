/*
Navicat MySQL Data Transfer

Source Server         : 1
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2021-06-05 16:49:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for zhenai_profile
-- ----------------------------
DROP TABLE IF EXISTS `zhenai_profile`;
CREATE TABLE `zhenai_profile` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) NOT NULL,
  `gender` varchar(10) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  `height` int(10) unsigned NOT NULL,
  `weight` int(10) unsigned NOT NULL,
  `income` varchar(255) NOT NULL,
  `mariage` varchar(50) NOT NULL,
  `education` varchar(50) NOT NULL,
  `xinzuo` varchar(50) NOT NULL,
  `house` varchar(50) NOT NULL,
  `car` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
