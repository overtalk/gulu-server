

CREATE TABLE `player` (
  `id` bigint(25) NOT NULL,
  `open_id` varchar(64) NOT NULL,
  `nickname` varchar(64) NOT NULL,
  `url` varchar(256) NOT NULL DEFAULT 'url',
  `sex` tinyint(4) NOT NULL DEFAULT 0,
  `strength` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `gold` bigint(20) NOT NULL DEFAULT 0,
  `diamond` bigint(20) NOT NULL DEFAULT 0,
  `experience` int(11) NOT NULL DEFAULT 0,
  `treasure` blob DEFAULT NULL,
  `daily` blob DEFAULT NULL,
  `achievement` blob DEFAULT NULL,
  `weapon` blob DEFAULT NULL,
  `weapon_equipped` blob DEFAULT NULL,
  `fashion` blob DEFAULT NULL,
  `fashion_equipped` blob DEFAULT NULL,
  `olditem` blob DEFAULT NULL,
  `pve` blob DEFAULT NULL,
  `arena` blob DEFAULT NULL,
  `last_pvp_room` text DEFAULT NULL,
  `last_arena` blob DEFAULT NULL,
  `free_box` blob DEFAULT NULL,
  `dailysign` blob DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL DEFAULT 0,
  `last_update_strength` int(10) unsigned NOT NULL DEFAULT 0,
  `invitation` blob DEFAULT NULL,
  `invitation_code` bigint(25) NOT NULL DEFAULT 0,
  `received_mail` blob DEFAULT NULL,
  `is_test` tinyint(1) DEFAULT 0,
  `free_lottery` blob DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `player_global_mail` (
  `player_id` bigint(25) NOT NULL,
  `mail_id` int(11) NOT NULL,
  `created_at` bigint(20) NOT NULL,
  PRIMARY KEY (`player_id`,`mail_id`),
  KEY `player_id` (`player_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `player_invitation` (
  `player_id` bigint(25) NOT NULL,
  `friend_id` bigint(25) NOT NULL,
  `created_at` int(10) NOT NULL,
  PRIMARY KEY (`player_id`,`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `pvp_result` (
  `player_id` bigint(25) NOT NULL,
  `detail` longblob DEFAULT NULL,
  `is_used` tinyint(1) DEFAULT 0,
  `last_ts` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `single_mail` (
  `id` bigint(25) NOT NULL AUTO_INCREMENT,
  `player_id` bigint(25) NOT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `icon` int(10) NOT NULL,
  `award` text DEFAULT NULL,
  `created_at` int(10) NOT NULL,
  `indate` int(10) NOT NULL,
  `image` text DEFAULT NULL,
  PRIMARY KEY (`id`,`player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000402 DEFAULT CHARSET=utf8mb4;

