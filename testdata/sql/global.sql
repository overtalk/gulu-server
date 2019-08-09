
CREATE TABLE `nodes` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `address` varchar(64) NOT NULL,
  `database` varchar(32) NOT NULL,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `maxconn` smallint(6) NOT NULL,
  `is_full` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0: idle, 1: full',
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`,`database`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `openid_mapping` (
  `player_id` bigint(25) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(64) NOT NULL,
  `node_id` smallint(6) DEFAULT '0',
  PRIMARY KEY (`player_id`),
  UNIQUE KEY `open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=587552 DEFAULT CHARSET=utf8mb4;

