-- 管理员表
CREATE TABLE IF NOT EXISTS `manager`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `username` VARCHAR(128) NOT NULL,
   `descript` VARCHAR(256) NOT NULL,
   `account`  VARCHAR(256) NOT NULL,    -- 账号(手机号)
   `password` VARCHAR(256) NOT NULL,    -- 密码
   `publickey` VARCHAR(256) NOT NULL,    -- 做那个登陆验证的
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO manager (username, descript, account, password, publickey)
VALUES ('pphui8', 'Ciallo～(∠・ω< )⌒★ ', '19861550668', '123212321', '3D75AD4DB3E952BC206E2DAFED2D91DC');

-- 用户表
CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `username` VARCHAR(128) NOT NULL,
   `account`  VARCHAR(256) NOT NULL,    -- 账号(手机号)
   `password` VARCHAR(256) NOT NULL,    -- 密码
   `publickey` VARCHAR(256) NOT NULL,    -- 做那个登陆验证的
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;