CREATE DATABASE IF NOT EXISTS cats;
USE cats;

CREATE TABLE user (
  `id`    INTEGER       NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name`  VARCHAR(255),
  `email` VARCHAR(1042) NOT NULL
);

CREATE TABLE podcast (
  `id`          INTEGER       NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title`       VARCHAR(1024),
  `image`       VARCHAR(2048),
  `lang`        VARCHAR(2),
  `category`    INTEGER,
  `description` VARCHAR(1000),
  `author`      VARCHAR(255)
);

CREATE TABLE episode (
  `id`          INTEGER       NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `podcast`     INTEGER       NOT NULL,
  `title`       VARCHAR(1024),
  `description` VARCHAR(1000),
  `guid`        VARCHAR(32),
  `pub_date`    TIMESTAMP     NOT NULL,
  `duration`    INTEGER,
  `explicit`    BOOLEAN
);

CREATE TABLE pivot_user_podcast (
  `user`    INTEGER NOT NULL PRIMARY KEY,
  `podcast` INTEGER NOT NULL
);

ALTER TABLE pivot_user_podcast
ADD CONSTRAINT
`FK_user_podcast` FOREIGN KEY (`podcast`) REFERENCES podcast (`id`);

ALTER TABLE episode
ADD CONSTRAINT
`FK_episode_podcast` FOREIGN KEY (`podcast`) REFERENCES podcast (`id`);
