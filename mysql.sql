-- 用户表
CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `username` VARCHAR(128) NOT NULL,
   `password` VARCHAR(128) NOT NULL,
   `descript` VARCHAR(256) NOT NULL,      -- 自我介绍
   `registerdate`  DATE NOT NULL,    -- 注册时间
   `phone` VARCHAR(256) NOT NULL,    -- 密码
   `privilege` VARCHAR(256) NOT NULL,    --  权限（普通用户、管理员
   `publickey` VARCHAR(256) NOT NULL,    -- 做那个登陆验证的
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('pphui8', '123212321', 'Ciallo～(∠・ω< )⌒★', '2022/8/30', '19861550668', 'manager', '72dce6e8801031109c58a89389bafb86');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('pphui89', '123212321', 'なんでやねん！', '2022/9/2', '19861550668', 'user', 'a7c42103ea7c1bc130033ff50f4333c9');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('User9527', '123212321', '9527号用户！', '2022/9/8', '19861550668', 'user', '137df37b15f276ee94f38c9fdf3d1d1c');

-- 商品表
CREATE TABLE IF NOT EXISTS `product`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `productname` VARCHAR(128) NOT NULL,
   `descript` VARCHAR(256) NOT NULL,      -- 介绍
   `price` INT UNSIGNED NOT NULL,    -- 价格
   `stock` INT UNSIGNED NOT NULL,    -- 库存
   `sale` INT UNSIGNED NOT NULL,    -- 销量
   `img` VARCHAR(256) NOT NULL,    -- 图片
   `type` VARCHAR(256) NOT NULL,    -- 类型
   `status` VARCHAR(256) NOT NULL,    -- 状态
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO product (productname, descript, price, stock, sale, img, type, status)
VALUES ('Nintendo Switch', '任天堂Nintendo Switch日版续航增强版，单机联机', 2199, 993, 7, 'https://img1.imgtp.com/2022/09/14/95WBnjLW.jpg', '硬件设备', '在售');

INSERT INTO product (productname, descript, price, stock, sale, img, type, status)
VALUES ('qq飞车游戏账号', 'qq飞车永久雷诺、包中A车、S车', 599, 1000, 0, 'https://img1.imgtp.com/2022/09/14/CVcWS5YZ.jpg', '游戏账号', '停售');

INSERT INTO product (productname, descript, price, stock, sale, img, type, status)
VALUES ('RTX3060', '华硕RTX3060、6GB、风冷游戏显卡', 3599, 399, 1, 'https://img1.imgtp.com/2022/09/14/Fq6dqGYA.jpg', '硬件设备', '在售');

INSERT INTO product (productname, descript, price, stock, sale, img, type, status)
VALUES ('iphone14 pro max', '8 + 256G土豪金配色，游戏体验极佳', 19999, 99, 1, 'https://img1.imgtp.com/2022/09/14/fPHpFjEa.jpg', '硬件设备', '在售');

-- 订单表
CREATE TABLE IF NOT EXISTS `orders`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `userid` INT UNSIGNED NOT NULL,    -- 购买的用户id
   `username` VARCHAR(128) NOT NULL,    -- 购买的用户
   `productid` INT UNSIGNED NOT NULL,   -- 购买的商品id
   `productname` VARCHAR(128) NOT NULL,   -- 购买的商品名称
   `price` INT UNSIGNED NOT NULL,    -- 购买的商品价格
   `number` INT UNSIGNED NOT NULL,    -- 购买的数量
   `date` DATE NOT NULL,    -- 购买时间
   `location` VARCHAR(256) NOT NULL,    -- 地址
   `status` VARCHAR(256) NOT NULL,    -- 订单状态
   PRIMARY KEY ( `id` ),
   FOREIGN KEY ( `userid` ) REFERENCES user ( `id` ),
   FOREIGN KEY ( `productid` ) REFERENCES product ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (1, 'pphui8', 2, 'qq飞车游戏账号', 599, 1, '2022/8/30', '上海', '已发货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (2, 'pphui89', 4, 'iphone14 pro max', 19999, 1, '2022/8/30', '云南', '已发货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (2, 'pphui89', 4, 'RTX3060', 3599, 1, '2022/9/15', '云南', '未付款');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (3, 'User9527', 4, 'iphone14 pro max', 19999, 1, '2022/9/15', '云南', '未发货');

CREATE TABLE IF NOT EXISTS `announcement`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `title` VARCHAR(128) NOT NULL,      -- 标题
   `content` VARCHAR(256) NOT NULL,      -- 内容
   `date` DATE NOT NULL,    -- 发布时间
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO announcement (title, content, date)
VALUES ('关于停服维护的通知', '本站将于9月20号进行为期2天的停服维护', '2022/9/15');