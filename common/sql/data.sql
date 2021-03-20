# 创建商品类
insert into element_class(name, picture_store_path,class_type) values("奶茶类","static/upload/温热1.jpeg",0);
insert into element_class(name, picture_store_path,class_type) values("水果茶类","static/upload/温热1.jpeg",0);
insert into element_class(name, picture_store_path,class_type) values("小吃类","static/upload/温热1.jpeg",0);
insert into element_class(name, picture_store_path,class_type) values("饮料类","static/upload/温热1.jpeg",0);

# 创建附属选项元素, 选项记录
insert into element(name, type, class_name, size, price, picture_store_path) values("温度",1,"附属选项类","冷饮",0,"static/upload/温热1.jpeg");
insert into element(name, type, class_name, size, price, picture_store_path) values("温度",1,"附属选项类","常温",0,"static/upload/温热1.jpeg");
insert into element(name, type, class_name, size, price, picture_store_path) values("温度",1,"附属选项类","热饮",0,"static/upload/温热1.jpeg");
insert into element_size_record(good_id,element_name,select_size) values(0,"温度","常温");

# 创建附属商品元素, 选项记录
insert into element(name, type, class_name, size, price, picture_store_path) values("珍珠",2,"附属商品类","少量",1,"static/upload/珍珠1.jpeg");
insert into element(name, type, class_name, size, price, picture_store_path) values("珍珠",2,"附属商品类","中量",2,"static/upload/珍珠2.jpeg");
insert into element(name, type, class_name, size, price, picture_store_path) values("珍珠",2,"附属商品类","大量",2.5,"static/upload/珍珠3.jpeg");
insert into element_size_record(good_id,element_name,select_size) values(0,"珍珠","少量");
