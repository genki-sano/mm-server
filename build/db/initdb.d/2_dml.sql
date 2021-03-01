START TRANSACTION;
INSERT INTO `users` (`type`,`name`) VALUES (0,'花子'),(1,'太郎');
INSERT INTO `categories` (`name`) VALUES ('食費'),('日用品'),('交通費'),('趣味'),('家具・家電'),('交際費'),('教養・教育'),('健康・医療'),('金融'),('住宅'),('水道・光熱費'),('通信費'),('税金'),('自動車');
COMMIT;
