满减功能，具体到订单
打折功能，具体到商品，还要有整单打折功能，还要考虑整单中有些商品是不能打折的…
打桌球是记时的，要和普通商品区分
麻将房是定时的..

问题
1. 订单号、Good号生成问题
2. 要考虑商品被删除后，订单记录的一致性
3. 要考虑规格..


功能
1. 商品的添加、商品选项的添加(如波霸奶茶有冰量、温度)
2. 商品订单的生成

流程
点单流程
1. 


约定:
1. Create 时，传入的是结构体
2. Update 时，通过函数名字控制字段
3. Get 时，通过函数名字控制字段
4. Del 时，通过函数名字控制字段
5. 日志字段在dao层详细输出

主品包括了 波霸奶茶(奶茶类)、水果茶(水果茶类)...    MainElement
附属商品 {  AttachElement          这个可以先自己sql写入
    调节选项 包括了 (少冰、多冰)(冰量)、(常温)(温度)...    
    丰富选项 包括了 (珍珠、红豆)(附属品)... 
}

这些东西归结为 Good
而 Good 指 MainGood + AttachGood

OrderGood 指 Good + 订单参数


订单>座位>物品


type Order struct{
    gorm.Model
    
    HasCheckout bool
    GoodPrice float64
    DeskPrice float64 // 考虑到之后价格可能更改
}

type Space struct {
    gorm.Model
    Name string // 无桌、桌球、普通桌
    Num int // 1号、2号...
    Price float64
    PriceRule int // 定时、记时
}

type ElementClass struct{
    name string;
    mainElementName string;
}

type Element struct {
	gorm.Model
	Name             string  `json:"name"`  
	Size string
	Price            string `json:"price"`
	PictureStorePath string  `json:"picture_store_path"`
	Type        ElementType     `json:"element_type"`   // 控制 main、attach
}

type MainElementAttachElementRecord struct {
    gorm.Model
    MainElementName string
    AttachElementName string
    DefaultSelectSize string
}
// ----
type OrderDeskRecord struct{
    gorm.Model
    OrderID int
    DeskID int
    StartTimestamp int
    EndTimestamp int
    PricePerHour float64
    FavorType int
}


type OrderMainGoodRecord struct{
    gorm.Model
    OrderID int
    MainGoodID int
    MainGoodSize string
    HasCheckout bool
    FavorType int
    MainGoodNum int
    DeskID int
    Price int   // 考虑到之后价格可能更改
}
// ----

type OrderRecord struct{
    gorm.Model
    OrderID int
    DeskID int
    MainGoodID int
    MainGoodSize string
    MainGoodNum int
    
    AttachGoodID int
    AttachGoodSize string
    Price int   // 考虑到之后价格可能更改
}


type ThingDao struct{
    gorm.Model
    GoodID int
}



与前端的交互结构为：Good、GoodClass

GoodClass: {
    Name string
    Goods []Good
}

Order: {
    ID int
    Desk Desk
    Favor Favor   // 订单折扣
}

Desk {
    Space Space
    startTimeStamp int
    endTimeStamp int
    Goods []Good
    Favor Favor   // 桌位折扣
    HadCheckOut bool
}

Favor {
    Name string
    Parameters []string
}



Good {
    MainElement Element
    AttachElements []Element
    
    
    HadCheckOut bool
    Favor Favor   // 物品折扣
    SelecDeskFavorType []FavorType
}

Space {
    Name string
    Num int
    Price float64
    PriceRule int
}


Element {
    Name string
    SelectableSizes []string
    SelectedIndexes []int
    Price float64
    PictureStorePath string
}


## 功能
1. 商品添加功能
    - 前端:
        - 选择商品图片(可空)
        - 填写商品名
        - 填写商品价格
        - 填写商品类
        - 填写附属选项(可空)
    
    - 后端:
        - 判断商品名(是否为空、是否存在...)
        - 判断商品图片(是否过大...)
        - 判断商品价格(是否为正数...)
        - 判断附属选项(是否存在...)
        - 在 Good 表中添加商品
        - (如果商品类不存在)在 GoodClass 表中添加商品类
        - 在 GoodAttachRecord 表中添加商品的附属商品记录(tip: 要记录默认选中选项)

2. 商品类获取功能
    - 前端:
    - 后端:
        - 在 GoodClass 表中获得所有类, 对每个类执行下面的操作
            - 在 Good 表中查找该类的所有 主品, 对于每个主品执行下面的操作
                - 在 GoodAttachRecord 表中查找商品附属选项，对每条记录执行下面的操作
                    - 形成 GoodAttachState，写入 GoodAttachStatus
               - 形成 GoodAttachStatus 后，将其写入 AttachClass
            -  形成 AttachClass 后，将其写入 Good
        - 形成 Good 后, 将其写入 GoodClass
          
3. 商品点单功能
    - 前端:
        - 选好商品，加入订单
        
    - 后端
        - 形成订单ID，写入 OrderLog 表
        - 对订单中的每个 OrderGood
            - 获取主商品ID
                - 让 orderID、mainGoodID 形成OrderMainGoodRecord,写入数据库
                - 对每个附属商品执行
                    - 获取附属商品ID，让 orderID、mainGoodID、attachGoodID 形成 OrderMainGoodAttachRecord，写入数据库
                
