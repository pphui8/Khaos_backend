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
VALUES ('new17860391537', '17860391535', '该用户还没有自我介绍', '2022/06/11', 'new518519', 'manager', '7f495cbb4a850873bd464392a282471e');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('new3701025', 'new7758521', '该用户还没有自我介绍', '2022/09/20', '17860391536', 'manager', '672d30ea79f6922e47c64a4e60dc7dd3');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('new5185191', 'new5185191', '该用户还没有自我介绍', '2022/8/30', '17860391538', 'manager', '05403d9154afe7b8a07c6fdddb321f12');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('坎公万岁', '7758521', '坎公真好玩', '2022/05/30', '17860391539', 'user', 'a7c42103ea7c1bc130033ff50f4333c9');

INSERT INTO user (username, password, descript, registerdate, phone, privilege, publickey)
VALUES ('pphui8', '123212321', 'Ciallo～(∠・ω< )⌒★', '2022/9/8', '19861550668', 'user', '137df37b15f276ee94f38c9fdf3d1d1c');

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

insert into `product` (productname, descript, price, stock, sale, img, type, status) 
values('BANPRESTO 龙珠Z 孙悟饭 景品手办', 'BANPRESTO出品手办', 119, 67, 2,'https://cbu01.alicdn.com/img/ibank/2019/904/262/12069262409_1782537758.jpg','Q版手办', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status) 
values('哔哩哔哩 环球萌动系列 小电视 晴雨伞', '小电视 晴雨伞', 68, 117, 3, 'https://img.alicdn.com/bao/uploaded/i2/2138521987/O1CN01bzjZs61QY5fy2bEmC_!!0-item_pic.jpg_300x300q90.jpg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('哔哩哔哩 全垒打系列 2233 毛巾', 'bilibili毛巾', 25, 211, 0,'https://p1.ssl.qhimgs1.com/sdr/400__/t01ba842aeaebc76bfb.jpg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('BILIBILIGOODS 软萌大礼包', 'bilibili软萌大礼包', 327, 111, 1, 'https://img13.360buyimg.com/n5/s800x800_jfs/t1/111266/4/5765/732846/5eb6456fE01bd1067/68fbf7f6ca6b2204.jpg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《DRAGON BALL超七龙珠超》1-17平装版 漫画 港台图书', '七龙珠图书', 425, 337, 0,'https://img.alicdn.com/bao/uploaded/i2/4217166954/O1CN01xICI0H21EyzfBOBjE_!!0-item_pic.jpg_310x310.jpg','漫画', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('万代 龙珠Z 弗利萨 终极形态 拼装模型', '万代拼装模型', 109, 71, 2,'https://p5.ssl.qhimgs1.com/sdr/400__/t011251f2df1f82e615.jpg','手办', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('BANPRSETO 龙珠Z 贝吉塔 景品手办', 'BANPRSETO景品手办', 119, 77, 1,'https://r3.hpoi.net.cn/gk/pic/s/2020/07/7f2e0b61c10644e18c4eb6fc1b69f6e6.jpg','手办', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('BANPRSETO 超级赛亚人之神 贝吉塔 景品手办', 'BANPRSETO景品手办', 45, 66, 0, 'https://p0.ssl.qhimgs1.com/sdr/400__/t0138df0d08b86a820a.jpg','手办', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《七龙珠彩色漫画 魔人布欧篇》1-6 平装版 漫画 港台图书', '七龙珠图书', 300, 118, 2, 'https://img.alicdn.com/bao/uploaded/i2/1793589415/O1CN01lbmtCk2JQ7ZkKXLPV_!!0-item_pic.jpg','漫画', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《DRAGON BALL 590 QUIZ BOOK七龙珠590解密打全》 平装版 漫画 港台图书', '七龙珠590解密打全图书', 40, 99, 0,'https://mhfm1cnc.cdndm5.com/38/37554/20170718134544_320x246_56.jpg','漫画', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《七龙珠彩色漫画 少年篇》1-8 平装版漫画 港台图书', '七龙珠漫画', 400, 87, 0, 'https://img.alicdn.com/bao/uploaded/i1/1793589415/O1CN01BmBrhd2JQ7Zm50XRr_!!0-item_pic.jpg','漫画', '停售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《七龙珠彩色漫画 人造人·赛鲁篇》1-6 平装版漫画 港台图书', '七龙珠漫画', 300, 71, 1, 'http://img3m4.ddimg.cn/62/22/1192492484-1_w_2.jpg','漫画', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('灵动创想 龙珠Z系列 钥匙扣滴胶挂件', '龙珠滴胶挂件', 29, 291, 1,'https://cbu01.alicdn.com/img/ibank/2019/168/243/10984342861_1384209607.jpg','摆件挂件', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('CALORIE 龙珠Z午睡套餐U型枕眼罩', '龙珠Z午睡套餐U型枕眼罩', 139, 44, 1, 'http://5b0988e595225.cdn.sohucs.com/images/20190701/11e74364c69140319191657af8096374.jpeg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('BANPRSETO（眼镜厂）龙珠GT 超级赛亚人之神4 贝吉塔 景品手办', '眼镜厂景品手办', 119, 77, 0, 'http://img.mp.itc.cn/upload/20170328/329c443003724b3a94aa20325eddb400_th.jpg','手办', '停售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('慢慢淘 龙珠系列 OAT恤 周边', '龙珠T恤', 99, 39, 0, 'https://img.alicdn.com/bao/uploaded/i4/2241650131/TB2UBommbJmpuFjSZFwXXaE4VXa_!!2241650131.jpg_300x300.jpg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('万代 S.H.Figuarts龙珠Z萨博 弗利萨军团 魂限定 可动手办', '可动手办', 429, 77, 2, 'http://n.sinaimg.cn/sinakd20100/728/w1170h1158/20211010/b5ff-94e950f9c64b5a8c245251f463babd30.jpg','手办', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('《七龙珠彩色漫画 GT版》1-8 平装版漫画 港台图书', '七龙珠漫画', 300, 21, 3, 'https://img.zcool.cn/community/01dfc25f31f832a801215aa038470c.jpg@2o.jpg','漫画', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('哔哩哔哩 小电视年糕毛绒抱枕 周边', 'bilibili抱枕周边', 159, 61, 3, 'https://gd1.alicdn.com/imgextra/i1/22898759/O1CN013tRADl2EZfuokj9UE_!!0-item_pic.jpg','日居用品', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('bilibili周边 哔哩哔哩小电视双层透明杯子 水杯 创意玻璃杯', 'bilibili周边创意玻璃杯', 80, 77, 1, 'https://img11.360buyimg.com/n5/s800x800_jfs/t1/195580/6/4261/67743/60a73135Efc8929f9/452c2d8f5418d79b.jpg','摆件挂件', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('bilibiliGoods 哔哩哔哩 小电视木质手机支架', 'bilibili手机支架', 38, 83, 2, 'https://img10.360buyimg.com/n5/s800x800_jfs/t1/194238/7/11134/102850/60dd850cE8adf09f3/40ed99f34f4e62df.jpg','摆件挂件', '在售');

insert into `product` (productname, descript, price, stock, sale, img, type, status)
values('bilibili哔哩哔哩 人生一串 小电视美食盒蛋', 'bilibili美食盒蛋', 59, 70, 3, 'https://r3.hpoi.net.cn/gk/cover/n/2019/09/74e973282d604464a1c07c70dbdc8377.jpg?date=1569403156','摆件挂件', '在售');

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

-- 28

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (1, 'new17860391537', 1, 'BANPRESTO 龙珠Z 孙悟饭 景品手办', 238, 2, '2022/9/25', '上海宝安', '已发货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (1, 'new17860391537', 6, '万代 龙珠Z 弗利萨 终极形态 拼装模型', 109, 1, '2022/9/25', '上海宝安', '已取消');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (2, 'new3701025', 3, '《七龙珠彩色漫画 魔人布欧篇》1-6 平装版 漫画 港台图书', 300, 1, '2022/9/25', '云南昆明', '未付款');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (3, 'new5185191', 3, '万代 S.H.Figuarts龙珠Z萨博 弗利萨军团 魂限定 可动手办', 429, 1, '2022/9/25', '山东烟台', '已收货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (2, 'new3701025', 20, 'bilibili周边 哔哩哔哩小电视双层透明杯子 水杯 创意玻璃杯', 80, 1, '2022/9/26', '云南昆明', '未付款');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (4, '渡鸦', 19, '哔哩哔哩 小电视年糕毛绒抱枕 周边', 122, 2, '2022/9/26', '北京大兴', '未付款');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (4, '渡鸦', 14, 'CALORIE 龙珠Z午睡套餐U型枕眼罩', 44, 1, '2022/9/27', '北京大兴', '已收货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (4, '渡鸦', 11, '《七龙珠彩色漫画 人造人·赛鲁篇》1-6 平装版漫画 港台图书', 87, 1, '2022/9/27', '北京大兴', '未发货');

INSERT INTO `orders` (userid, username, productid, productname, price, number, date, location, status)
VALUES (3, 'new5185191', 22, 'bilibili哔哩哔哩 人生一串 小电视美食盒蛋', 59, 1, '2022/9/27', '山东烟台', '未付款');

CREATE TABLE IF NOT EXISTS `announcement`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `title` VARCHAR(128) NOT NULL,      -- 标题
   `content` VARCHAR(256) NOT NULL,      -- 内容
   `date` DATE NOT NULL,    -- 发布时间
   `img` VARCHAR(256) NOT NULL,    -- 图片(可为空)
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into announcement (title, content, date, img) 
values('【9/7】9月7日公告更新', '您好，为了进一步提高游戏运行质量，兹定于2022年9月8日(周三)09：30开始例行维护，预计维护1个小时，维护期间所有游戏房间将无法进入，请大家在维护结束之后再进行游戏，给您带来的不便敬请谅解！  　　xxxx游戏中心，提倡健康、休闲、公平的游戏理念。游戏账号设有充值限额，游戏消耗封顶限额，同时平台还提供了大量的免费游戏及比赛。维护期间将进行网络优化，如果您在更新过程遇到问题，请及时联系客服解决。  　　如果您发现任何BUG或想提出建议，请通过在线客服或反映，将有机会获得一定的奖励！', '2022/09/05', '');

insert into announcement (title, content, date, img) 
values('【9/12】9月12日公告更新', '2022年9月17日本周一11:00-12:00，龙珠官方将对页面商城进行临时维护。  　　届时将下架2022年9月9日15：00~10月17日10：00七天乐【活动五：商城大促销】活动的各个优惠礼包。仍有需要的玩家，请抓紧时间购买商城物品。  　　因维护给您带来的不便敬请谅解，感谢您长久以来对我们工作的理解与支持', '2022/09/11', '');

insert into announcement (title, content, date, img) 
values('【9/21】9月21日公告更新', '定于2022年9月22日15：00—18：00进行停服更新维护，请玩家相互转告。  　　停服时间：9月22日15：00—18：00  　　停服范围：全区全服  　　更新内容：修复游戏内部分BUG', '2022-09-21', 'https://img1.imgtp.com/2022/09/21/L0n3i5bw.jpg');

insert into announcement (title, content, date, img) 
values('【9/2】紧急公告', '开发人员要累死了，求打分多给两分，这个项目一个人的工作量这些应该还可以吧', '2022/09/02', '');


CREATE TABLE IF NOT EXISTS `post`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `userid` INT UNSIGNED NOT NULL,      -- 用户id
   `username` VARCHAR(128) NOT NULL,      -- 用户名
   `title` VARCHAR(128) NOT NULL,      -- 标题
   `content` MEDIUMTEXT NOT NULL,      -- 内容
   `browseNumber` INT UNSIGNED NOT NULL,    -- 浏览量
   `date` DATE NOT NULL,    -- 发布时间
   `legal` TINYINT NOT NULL, -- 是否合法
   `elite` TINYINT NOT NULL, -- 是否加精
   `img` VARCHAR(256) NOT NULL,    -- 图片(可为空)
   `tag` VARCHAR(256) NOT NULL,    -- 标签
   PRIMARY KEY ( `id` ),
   FOREIGN KEY ( `userid` ) REFERENCES user ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('4', '渡鸦', '我想睡觉', '累死了救命', '350', '2022-08-30', '1', '1','','日常');

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('4','渡鸦','我好饿啊','咋还不下课啊，老师您不饿吗','250','2022-09-15', '1', '1','https://p3.ssl.qhimgs1.com/t01197ecd9f3154a7f9.jpg','');

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('4','渡鸦','靠背咧卡成狗','啊不是这游戏优化简直有毒','117','2022-09-14','1','1','https://image.so.com/view?q=%E7%86%8A%E7%8C%AB%E8%A1%A8%E6%83%85%E5%8C%85&src=tab_www&correct=%E7%86%8A%E7%8C%AB%E8%A1%A8%E6%83%85%E5%8Chttp://www.yutudou.com/uploads/allimg/170524/1-1F5241Q938.jpg','日常');

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('3','坎公万岁','饿饿饿','想吃章鱼小丸子','75','2022-09-11','1','0','http://n.sinaimg.cn/translate/20170929/Szxw-fymkwyr9687731.jpg', '日常');

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('3','new5185191','为什么我抽不到小公主啊啊啊啊啊！！！！','我愿意用老左一辈子单身换个小公主（狗头保命）','60','2022-08-09','1','1','https://gd-hbimg.huaban.com/6f591e93773996703519beed13cb875a1569699a44e7-B85Uah_fw236','同人');

insert into `post` (`userid`, `username`, `title`, `content`, `browseNumber`, `date`, `legal`, `elite`, `img`, `tag`)
values('4','渡鸦','好多作业','救命颈椎要断了','45','2022-09-13','1','0','','攻略');


CREATE TABLE IF NOT EXISTS `comment`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `userid` INT UNSIGNED NOT NULL,      -- 用户id
   `username` VARCHAR(128) NOT NULL,      -- 用户名
   `postid` INT UNSIGNED NOT NULL,      -- 帖子id
   `content` MEDIUMTEXT NOT NULL,      -- 内容
   `date` DATE NOT NULL,    -- 发布时间
   `support` INT UNSIGNED NOT NULL,    -- 点赞数
   `against` INT UNSIGNED NOT NULL,    -- 点踩数
   PRIMARY KEY ( `id` ),
   FOREIGN KEY ( `userid` ) REFERENCES user ( `id` ),
   FOREIGN KEY ( `postid` ) REFERENCES post ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (1, 'new17860391537', 5, '原来还有人没有抽到大公主吗笑死', '2022/05/26', 24, 4);

INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (1, 'new17860391537', 5, '楼主老倒霉蛋了笑死', '2022/09/24', 65, 0);

INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (1, 'new17860391537', 5, '楼主老倒霉蛋了笑死', '2022/09/24', 1, 0);

INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (1, 'new17860391537', 5, '一首凉凉送给你', '2022/09/24', 0, 1);

INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (5, 'pphui8', 1, '睡吧，吃完就睡小心胖成猪', '2022/05/26', 2, 0);

INSERT INTO comment (userid, username, postid, content, date, support, against)
VALUES (5, 'pphui8', 3, '怕不是服务器炸了（物理）', '2022/05/24', 2, 0);