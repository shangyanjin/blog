/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 100509 (10.5.9-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 100509 (10.5.9-MariaDB)
 File Encoding         : 65001

 Date: 23/02/2025 22:43:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_admin
-- ----------------------------
DROP TABLE IF EXISTS `blog_admin`;
CREATE TABLE `blog_admin`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Account',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Password',
  `salt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Password Salt',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'User Name',
  `first_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'First Name',
  `last_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Last Name',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Avatar',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Title',
  `about` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'About',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Mobile',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Phone',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Email',
  `twitter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Twitter',
  `facebook` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Facebook',
  `linkedin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'LinkedIn',
  `id_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ID Card Number',
  `gender` tinyint NULL DEFAULT 0 COMMENT 'Gender: 0=Unknown,1=Male,2=Female',
  `birthday` date NULL DEFAULT NULL COMMENT 'Birthday',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber',
  `type` int NULL DEFAULT 0 COMMENT 'User Type: 0=Default,1=User,2=Creator',
  `terms_accepted` int NULL DEFAULT NULL COMMENT 'Terms Accepted: 0=Default,1=Accepted,2=Declined',
  `newsletter` int NULL DEFAULT NULL COMMENT 'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed',
  `post` int NULL DEFAULT 0 COMMENT 'Post Count',
  `level` int NULL DEFAULT 0 COMMENT 'Level',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT 'Last Login Time',
  `last_login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Last Login IP',
  `created_at` timestamp NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_account`(`account` ASC) USING BTREE,
  INDEX `idx_sort`(`sort` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Admin Users Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_admin
-- ----------------------------

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int NOT NULL DEFAULT 0 COMMENT 'Parent Category ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Category Name',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Category Icon',
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Category Slug',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Category Description',
  `meta_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO Title',
  `meta_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO Description',
  `post_count` int NOT NULL DEFAULT 0 COMMENT 'Post Count',
  `level` int NULL DEFAULT 0 COMMENT 'Level',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Disabled,1=Active',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_parent`(`parent_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Categories Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_category
-- ----------------------------

-- ----------------------------
-- Table structure for blog_channel
-- ----------------------------
DROP TABLE IF EXISTS `blog_channel`;
CREATE TABLE `blog_channel`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Channel ID',
  `user_id` int NOT NULL DEFAULT 0 COMMENT 'User ID',
  `creator_id` int NOT NULL DEFAULT 0 COMMENT 'Creator ID',
  `store_id` int NOT NULL DEFAULT 0 COMMENT 'Store ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Channel Name',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'Channel Description',
  `sort` int NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Active,2=Disabled,3=Private,4=Muted',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Media Channels Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_channel
-- ----------------------------

-- ----------------------------
-- Table structure for blog_collection
-- ----------------------------
DROP TABLE IF EXISTS `blog_collection`;
CREATE TABLE `blog_collection`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int NOT NULL COMMENT 'User ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Name',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Title',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Icon',
  `resource` int NULL DEFAULT NULL COMMENT 'Resource Count',
  `download` int NOT NULL COMMENT 'Download Count',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Published,2=Draft',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Collections Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_collection
-- ----------------------------

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NULL DEFAULT NULL COMMENT 'User ID',
  `post_id` int NULL DEFAULT NULL COMMENT 'Post ID',
  `parent_id` int NULL DEFAULT NULL COMMENT 'Parent Comment ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Post Title',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Comment Content',
  `pic` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Image URL',
  `list_pic` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Image List',
  `video` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Video',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Name',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Avatar',
  `likes` int NULL DEFAULT 0 COMMENT 'Likes Count',
  `dislikes` int NULL DEFAULT 0 COMMENT 'Dislikes Count',
  `ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'IP Address',
  `is_anonymous` varchar(4) NULL DEFAULT '0' COMMENT 'Is Anonymous: 0=No,1=Yes',
  `is_top` varchar(4) NULL DEFAULT '0' COMMENT 'Is Top: 0=No,1=Yes',
  `is_hot` varchar(4) NULL DEFAULT '0' COMMENT 'Is Hot: 0=No,1=Yes',
  `is_hidden` varchar(4) NULL DEFAULT '0' COMMENT 'Is Hidden: 0=No,1=Yes',
  `log` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Log',
  `sort` int NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Active,2=Hidden,3=Deleted',
  `created_at` datetime NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` datetime NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deleted Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `post_id`(`post_id` ASC) USING BTREE,
  INDEX `parent_id`(`parent_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 155 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Comments Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_comment
-- ----------------------------

-- ----------------------------
-- Table structure for blog_download
-- ----------------------------
DROP TABLE IF EXISTS `blog_download`;
CREATE TABLE `blog_download`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NOT NULL COMMENT 'User ID',
  `resource_id` int NOT NULL COMMENT 'Resource ID',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_date`(`user_id` ASC, `created_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Download Records Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_download
-- ----------------------------

-- ----------------------------
-- Table structure for blog_follow
-- ----------------------------
DROP TABLE IF EXISTS `blog_follow`;
CREATE TABLE `blog_follow`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NOT NULL COMMENT 'User ID',
  `followed_id` int NOT NULL COMMENT 'Followed ID',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_to`(`user_id` ASC, `followed_id` ASC) USING BTREE,
  INDEX `idx_to`(`followed_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Follow Records Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_follow
-- ----------------------------

-- ----------------------------
-- Table structure for blog_level
-- ----------------------------
DROP TABLE IF EXISTS `blog_level`;
CREATE TABLE `blog_level`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Level Name',
  `daily_limit` int NOT NULL DEFAULT 0 COMMENT 'Daily Download Limit',
  `weekly_limit` int NOT NULL DEFAULT 0 COMMENT 'Weekly Download Limit',
  `monthly_limit` int NOT NULL DEFAULT 0 COMMENT 'Monthly Download Limit',
  `yearly_limit` int NOT NULL DEFAULT 0 COMMENT 'Yearly Download Limit',
  `total_limit` int NOT NULL DEFAULT 0 COMMENT 'Total Download Limit (0=Unlimited)',
  `price` decimal(10, 2) NOT NULL COMMENT 'Price',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Disabled,1=Active',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'User Levels Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_level
-- ----------------------------

-- ----------------------------
-- Table structure for blog_like
-- ----------------------------
DROP TABLE IF EXISTS `blog_like`;
CREATE TABLE `blog_like`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NOT NULL COMMENT 'User ID',
  `to_id` int NOT NULL COMMENT 'Target ID',
  `type` tinyint(1) NOT NULL COMMENT 'Type: 1=Resource,2=Creator',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_to`(`user_id` ASC, `to_id` ASC, `type` ASC) USING BTREE,
  INDEX `idx_to`(`to_id` ASC, `type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Like Records Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_like
-- ----------------------------

-- ----------------------------
-- Table structure for blog_log
-- ----------------------------
DROP TABLE IF EXISTS `blog_log`;
CREATE TABLE `blog_log`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `merchant_id` int NULL DEFAULT 0 COMMENT 'Merchant ID',
  `store_id` int UNSIGNED NULL DEFAULT 0 COMMENT 'Store ID',
  `user_id` int NULL DEFAULT NULL COMMENT 'User ID',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'User Name',
  `type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Operation Type',
  `channel` int UNSIGNED NULL DEFAULT 0 COMMENT 'Channel Number',
  `operator_id` int NULL DEFAULT NULL COMMENT 'Operator ID',
  `operator_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Operator Name',
  `amount` decimal(12, 2) NULL DEFAULT NULL COMMENT 'Amount',
  `score` int NULL DEFAULT 0 COMMENT 'Score',
  `level` int NULL DEFAULT 0 COMMENT 'Level',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Log Content',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `os` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Operating System',
  `ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'IP Address',
  `browser` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Browser',
  `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Created Time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 183 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'System Logs Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_log
-- ----------------------------

-- ----------------------------
-- Table structure for blog_message
-- ----------------------------
DROP TABLE IF EXISTS `blog_message`;
CREATE TABLE `blog_message`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NOT NULL COMMENT 'User ID',
  `ticket_id` int NOT NULL DEFAULT 0 COMMENT 'Ticket Number',
  `parent_id` int NOT NULL DEFAULT 0 COMMENT 'Parent ID',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Role: 0=Default,1=User,2=Creator,3=Platform',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Avatar',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other',
  `priority` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Priority: 0=Low,1=Medium,2=High',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Message Title',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Content',
  `pic` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Images(Comma Separated)',
  `attachment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Attachment',
  `assign` int NULL DEFAULT NULL COMMENT 'Assigned To',
  `note` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Internal Note',
  `log` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'System Log',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Pending,1=Processing,2=Completed,3=Closed',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_parent`(`parent_id` ASC) USING BTREE,
  INDEX `idx_user`(`user_id` ASC) USING BTREE,
  INDEX `idx_user_type`(`role` ASC) USING BTREE,
  INDEX `idx_type`(`type` ASC) USING BTREE,
  INDEX `idx_priority`(`priority` ASC) USING BTREE,
  INDEX `idx_status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Messages Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_message
-- ----------------------------

-- ----------------------------
-- Table structure for blog_page
-- ----------------------------
DROP TABLE IF EXISTS `blog_page`;
CREATE TABLE `blog_page`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NULL DEFAULT NULL COMMENT 'User ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Title',
  `summary` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Summary',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Content',
  `meta_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO Title',
  `meta_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO Keywords',
  `meta_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'SEO Description',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Published,2=Draft',
  `created_at` datetime NULL DEFAULT NULL COMMENT 'Created Time',
  `updated_at` datetime NULL DEFAULT NULL COMMENT 'Updated Time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deleted Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Pages Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_page
-- ----------------------------

-- ----------------------------
-- Table structure for blog_post
-- ----------------------------
DROP TABLE IF EXISTS `blog_post`;
CREATE TABLE `blog_post`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int NULL DEFAULT NULL COMMENT 'User ID',
  `category_id` int NULL DEFAULT NULL COMMENT 'Category ID',
  `category_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Category Name',
  `type` int NULL DEFAULT 0 COMMENT 'Type: 0=Default,1=Article,2=Video,3=Audio,4=Resource',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Title',
  `summary` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Post Summary or Excerpt',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Content',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Author',
  `revenue` decimal(10, 2) NULL DEFAULT 0.00 COMMENT 'Revenue',
  `price` decimal(10, 2) NULL DEFAULT 0.00 COMMENT 'Price',
  `keyword` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Keywords',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Summary',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Tags',
  `pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Image URL',
  `list_pic` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Image List',
  `video` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Video',
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Cover Image',
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'File Name',
  `file_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'File URL',
  `file_size` bigint NULL DEFAULT 0 COMMENT 'File Size',
  `file_md5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'File Hash',
  `rating` decimal(10, 2) NULL DEFAULT 0.00 COMMENT 'Rating',
  `duration` int NULL DEFAULT 0 COMMENT 'Duration(seconds)',
  `ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'IP Address',
  `likes` int NULL DEFAULT 0 COMMENT 'Likes Count',
  `dislikes` int NULL DEFAULT 0 COMMENT 'Dislikes Count',
  `views` int NULL DEFAULT 0 COMMENT 'Views Count',
  `downloads` int NULL DEFAULT 0 COMMENT 'Downloads Count',
  `collects` int NULL DEFAULT 0 COMMENT 'Collections Count',
  `comments` int NULL DEFAULT 0 COMMENT 'Comments Count',
  `is_new` varchar(4) NULL DEFAULT '0' COMMENT 'Is New: 0=No,1=Yes',
  `is_hot` varchar(4) NULL DEFAULT '0' COMMENT 'Is Hot: 0=No,1=Yes',
  `is_recommend` varchar(4) NULL DEFAULT '0' COMMENT 'Is Recommended: 0=No,1=Yes',
  `is_top` varchar(4) NULL DEFAULT '0' COMMENT 'Is Top: 0=No,1=Yes',
  `is_free` varchar(4) NULL DEFAULT '0' COMMENT 'Is Free: 0=Default,1=Free,2=Member,3=Paid',
  `is_review` varchar(4) NULL DEFAULT '0' COMMENT 'Review: 0=Default,1=Normal,2=No Comments,3=Comments Visible to Author',
  `sort` int NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Active,2=Stopped,3=Under Review',
  `created_at` datetime NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` datetime NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `category_id`(`category_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 155 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Posts Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_post
-- ----------------------------

-- ----------------------------
-- Table structure for blog_region
-- ----------------------------
DROP TABLE IF EXISTS `blog_region`;
CREATE TABLE `blog_region`  (
  `id` smallint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `parent_id` smallint UNSIGNED NOT NULL COMMENT 'Parent ID',
  `name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Name',
  `type` tinyint(1) NOT NULL DEFAULT 2 COMMENT 'Type',
  `sort` int NULL DEFAULT 0 COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Disabled,1=Active',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `parent_id`(`parent_id` ASC) USING BTREE,
  INDEX `region_type`(`type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4044 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Regions Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_region
-- ----------------------------

-- ----------------------------
-- Table structure for blog_setting
-- ----------------------------
DROP TABLE IF EXISTS `blog_setting`;
CREATE TABLE `blog_setting`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Setting Name',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Setting Value',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Settings Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_setting
-- ----------------------------

-- ----------------------------
-- Table structure for blog_site
-- ----------------------------
DROP TABLE IF EXISTS `blog_site`;
CREATE TABLE `blog_site`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT 0 COMMENT 'User ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Site Name',
  `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Domain',
  `tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Telephone',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Mobile Phone',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Email',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Enterprise Website System' COMMENT 'Site Title',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Professional Enterprise Website System' COMMENT 'Site Description',
  `keyword` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Enterprise Website,Website System' COMMENT 'Keywords',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Beijing Chaoyang District' COMMENT 'Address',
  `contact` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Manager Zhang' COMMENT 'Contact Person',
  `fax` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Fax',
  `qq` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '88888888' COMMENT 'QQ',
  `wechat` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'company_wechat' COMMENT 'WeChat',
  `icp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'ICP 12345678' COMMENT 'ICP License',
  `mit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'MIT License' COMMENT 'MIT License',
  `police` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Police Record 123456' COMMENT 'Police Record',
  `privacy` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Privacy Policy',
  `service` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Terms of Service',
  `user` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'User Agreement',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Agent Agreement',
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '/static/logo.png' COMMENT 'Logo',
  `favicon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '/static/favicon.ico' COMMENT 'Favicon',
  `banner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '/static/banner.jpg' COMMENT 'Banner Image',
  `footer` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Footer Info',
  `copyright` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '©2024 All Rights Reserved' COMMENT 'Copyright Info',
  `code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Statistics Code',
  `seo_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Enterprise Website System' COMMENT 'SEO Title',
  `seo_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Professional Enterprise Website System' COMMENT 'SEO Description',
  `seo_keyword` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Enterprise Website,Website System' COMMENT 'SEO Keywords',
  `maintenance` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Maintenance Notice' COMMENT 'Maintenance Notice',
  `theme` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'default' COMMENT 'Theme',
  `language` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'en-US' COMMENT 'Language',
  `company` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'XX Technology Co., Ltd.' COMMENT 'Company Name',
  `pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'Enterprise System' COMMENT 'Picture',
  `static` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '/static/logo.png' COMMENT 'Static Files',
  `sort` int NULL DEFAULT NULL COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Active,2=Closed',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT 'Created Time',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Site Configuration Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_site
-- ----------------------------

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `user_id` int UNSIGNED NULL DEFAULT 0 COMMENT 'User ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Name',
  `sort` int NULL DEFAULT NULL COMMENT 'Sort Order',
  `status` varchar(4) NULL DEFAULT '0' COMMENT 'Status: 0=Default,1=Active,2=Disabled',
  `created_at` datetime NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` datetime NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT 'Deleted Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Article Tags Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------

-- ----------------------------
-- Table structure for blog_ticket
-- ----------------------------
DROP TABLE IF EXISTS `blog_ticket`;
CREATE TABLE `blog_ticket`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `parent_id` int NOT NULL DEFAULT 0 COMMENT 'Parent ID',
  `user_id` int NOT NULL COMMENT 'User ID',
  `ticket_id` int NOT NULL DEFAULT 0 COMMENT 'Ticket Number',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Role: 0=Default,1=User,2=Creator,3=Platform',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'Avatar',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Type: 0=Feedback,1=Account,2=Payment,3=Suggestion,4=Business,5=Other',
  `priority` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Priority: 0=Low,1=Medium,2=High',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Ticket Title',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Content',
  `pic` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Images(Comma Separated)',
  `attachment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Attachment',
  `assign` int NULL DEFAULT NULL COMMENT 'Assigned To',
  `note` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'Internal Note',
  `log` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'System Log',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Pending,1=Processing,2=Completed,3=Closed',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_parent`(`parent_id` ASC) USING BTREE,
  INDEX `idx_user`(`user_id` ASC) USING BTREE,
  INDEX `idx_user_type`(`role` ASC) USING BTREE,
  INDEX `idx_type`(`type` ASC) USING BTREE,
  INDEX `idx_priority`(`priority` ASC) USING BTREE,
  INDEX `idx_status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Support Tickets Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_ticket
-- ----------------------------

-- ----------------------------
-- Table structure for blog_upload
-- ----------------------------
DROP TABLE IF EXISTS `blog_upload`;
CREATE TABLE `blog_upload`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `cid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Category ID',
  `uid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'User ID',
  `type` tinyint UNSIGNED NOT NULL DEFAULT 10 COMMENT 'File Type: [10=Image, 20=Video]',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'File Name',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'File Path',
  `ext` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'File Extension',
  `size` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'File Size',
  `is_delete` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Is Deleted: 0=No,1=Yes',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Created Time',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Updated Time',
  `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Deleted Time',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cid`(`cid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1537 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Uploaded Files Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_upload
-- ----------------------------

-- ----------------------------
-- Table structure for blog_upload_cate
-- ----------------------------
DROP TABLE IF EXISTS `blog_upload_cate`;
CREATE TABLE `blog_upload_cate`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Parent ID',
  `type` tinyint UNSIGNED NOT NULL DEFAULT 10 COMMENT 'Type: [10=Image, 20=Video]',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Category Name',
  `is_delete` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Is Deleted: 0=No,1=Yes',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Created Time',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Updated Time',
  `deleted_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'Deleted Time',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Upload Categories Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_upload_cate
-- ----------------------------

-- ----------------------------
-- Table structure for blog_user
-- ----------------------------
DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Account',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'Password',
  `salt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Password Salt',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'User Name',
  `first_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'First Name',
  `last_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Last Name',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Avatar',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Title',
  `about` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'About',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Mobile',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Phone',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Email',
  `twitter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Twitter',
  `facebook` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Facebook',
  `linkedin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'LinkedIn',
  `id_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ID Card Number',
  `gender` tinyint NULL DEFAULT 0 COMMENT 'Gender: 0=Unknown,1=Male,2=Female',
  `birthday` date NULL DEFAULT NULL COMMENT 'Birthday',
  `role` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Role: 0=Default,1=Admin,2=Manager,3=Editor,4=Author,5=Subscriber',
  `type` int NULL DEFAULT 0 COMMENT 'User Type: 0=Default,1=User,2=Creator',
  `terms_accepted` int NULL DEFAULT NULL COMMENT 'Terms Accepted: 0=Default,1=Accepted,2=Declined',
  `newsletter` int NULL DEFAULT NULL COMMENT 'Newsletter: 0=Default,1=Subscribed,2=Unsubscribed',
  `post` int NULL DEFAULT 0 COMMENT 'Post Count',
  `level` int NULL DEFAULT 0 COMMENT 'Level',
  `status` varchar(4) NOT NULL DEFAULT '0' COMMENT 'Status: 0=Disabled,1=Active,2=Unverified,3=Deleted,4=Frozen,5=Pending',
  `sort` int NOT NULL DEFAULT 0 COMMENT 'Sort Order',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT 'Last Login Time',
  `last_login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Last Login IP',
  `created_at` timestamp NULL DEFAULT current_timestamp COMMENT 'Created Time',
  `updated_at` timestamp NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_account`(`account` ASC) USING BTREE,
  INDEX `idx_sort`(`sort` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'Users Table' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
