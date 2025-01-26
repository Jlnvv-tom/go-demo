CREATE TABLE user (
  id bigint AUTO_INCREMENT,
  name varchar(255) NULL COMMENT 'The username',
  phone varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
  UNIQUE name_index (name),
  PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';