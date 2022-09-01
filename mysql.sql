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
VALUES ('pphui8', '123212321', 'Ciallo～(∠・ω< )⌒★', '2022/8/30', '19861550668', 'manager', '3D75AD4DB3E952BC206E2DAFED2D91DC');

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
VALUES ('伸腿瞪眼丸', '药到病除', 999, 993, 7, 'none', '保健品', '在售');

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
VALUES (1, 'pphui8', 1, '伸腿瞪眼丸', 999, 1, '2022/8/30', '上海', '已发货');