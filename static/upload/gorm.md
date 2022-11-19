# SQL 构建器



# 插入

## 用指定的字段创建记录

创建记录并更新给出的字段。

```go
user := &User{
   Name: "氧气语句",
   Age:  200,
}
db.Debug().Select("Age").Create(&user)
// INSERT INTO `users` (`age`) VALUES (200)
```

创建一个记录且一同忽略传递给略去的字段值。

```go
user := &User{
   Name: "氧气语句",
   Age:  200,
}
db.Debug().Omit("name").Create(&user)
//  INSERT INTO `users` (`age`) VALUES (200)
```

## 批量插入



```go
var users = []User{
   {Name: "氧气爸爸", Age: 12},
   {Name: "氧气爸爸2", Age: 13},
   {Name: "氧气爸爸3", Age: 3},
   {Name: "氧气爸爸4", Age: 19},
}
db.Debug().Create(&users)
// INSERT INTO `users` (`name`,`age`) VALUES ('氧气爸爸',12),('氧气爸爸2',13),('氧气爸爸3',3),('氧气爸爸4',19)
for _, user := range users {
   fmt.Println(user.ID)
}
```

使用 `CreateInBatches` 分批创建时，你可以指定每批的数量，例如：

```go
var users = []User{
   {Name: "氧气爸爸", Age: 12},
   {Name: "氧气爸爸2", Age: 13},
   {Name: "氧气爸爸3", Age: 3},
   {Name: "氧气爸爸4", Age: 19},
}
db.Debug().CreateInBatches(users, 2)
// 每一次执行两条数据插入
// INSERT INTO `users` (`name`,`age`) VALUES ('氧气爸爸',12),('氧气爸爸2',13)
//INSERT INTO `users` (`name`,`age`) VALUES ('氧气爸爸3',3),('氧气爸爸4',19)
for _, user := range users {
   fmt.Println(user.ID)
}
```

## 创建钩子

GORM 允许用户定义的钩子有 `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate` 创建记录时将调用这些钩子方法，请参考 [Hooks](https://gorm.io/zh_CN/docs/hooks.html) 中关于生命周期的详细信息

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```

如果您想跳过 `钩子` 方法，您可以使用 `SkipHooks` 会话模式，例如：

```go
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
```



## 根据 Map 创建

GORM 支持根据 `map[string]interface{}` 和 `[]map[string]interface{}{}` 创建记录，例如：

```go
db.Model(&User{}).Create(map[string]interface{}{
   "Name": "语音YY", "Age": 18,
})
db.Model(&User{}).Create([]map[string]interface{}{
   {"Name": "金祖", "Age": 19},
   {"Name": "金祖", "Age": 19},
})
```

## 高级选项

### 关联创建

创建关联数据时，如果关联值是非零值，这些关联会被 upsert，且它们的 `Hook` 方法也会被调用

```go

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

db.AutoMigrate(&CreditCard{})
db.AutoMigrate(&User{})
db.Create(&User{
   Name: "氧气语音",
   CreditCard: CreditCard{
      Number: "44232222222222222",
   },
})
	// INSERT INTO `credit_cards`
	//INSERT INTO `users`
```

# 查询



## Or 条件

```go
log.Println("=======================Or 条件")
var settings []ImSetting
db.Where("value = ?", "机制").Or("value = ?", "神丹").Find(&settings)
log.Println(settings)
// SELECT * FROM users WHERE value = '机制' OR value = '神丹';

var settings2 []ImSetting
// Struct
db.Where("project_name = '焦糖'").Or(ImSetting{ProjectName: "氧气", BigType: "90"}).Find(&settings2)
// SELECT * FROM users WHERE project_name = '焦糖' OR (project_name = '氧气 2' AND bigtype = 90);
log.Println(settings2)

var settings3 []ImSetting
// Map
db.Where("project_name = '焦糖'").Or(map[string]interface{}{"project_name": "Miya2", "bigtype": "80"}).Find(&settings3)
// SELECT * FROM users WHERE project_name = '焦糖' OR (project_name = 'Miya2' AND bigtype = 80);
log.Println(settings3)
```

## Select

```go
var settings []ImSetting
db.Debug().Select("project_name,value").Find(&settings)
// SELECT name, age FROM users;
log.Println(settings)

var settings2 []ImSetting
db.Select([]string{"project_name", "value"}).Find(&settings2)
// SELECT name, age FROM users;
log.Println(settings2)
var settings3 []ImSetting
db.Table("im_setting").Select("COALESCE(bigtype,?)", "90").Rows()
// SELECT COALESCE(age,'42') FROM users;
log.Println(settings3)
```

## 排序

```go
var settings []ImSetting
db.Order("id asc, project_name").Find(&settings)
// SELECT * FROM users ORDER BY age desc, name;
log.Println(settings)
// Multiple orders
var settings2 []ImSetting
db.Select("bigtype, project_name").Order("bigtype asc").Order("project_name").Find(&settings2)
// SELECT * FROM users ORDER BY age desc, name;
log.Println(settings2)
```

## 分页（Limit & Offset）

`Limit`指定要检索的最大记录数。 `Offset`指定在开始返回记录前要跳过的记录数。

```go
var settings []ImSetting
db.Limit(3).Find(&settings)
// SELECT * FROM users LIMIT 3;
log.Println(settings)

var settings2 []ImSetting
var settings3 []ImSetting
// Cancel limit condition with -1
db.Debug().Limit(2).Find(&settings2).Limit(1).Find(&settings3)
//同时执行了两条的sql  语句
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)
log.Println(settings2)
log.Println(settings3)

log.Println("========================Offset=")
var settings4 []ImSetting
db.Offset(10).Find(&settings4)
// SELECT * FROM `im_setting`
log.Println(settings4)

var settings5 []ImSetting
db.Debug().Limit(2).Offset(1).Find(&settings5)
//在第一条数据后面取2条数据的   也就是2，3
// SELECT * FROM `im_setting`   LIMIT 2 OFFSET 1
log.Println(settings5)

var settings6 []ImSetting
var settings7 []ImSetting
// Cancel offset condition with -1
db.Debug().Offset(10).Find(&settings6).Offset(1).Find(&settings7)
// SELECT * FROM `im_setting` s
//  SELECT * FROM `im_setting`
log.Println(settings6)
log.Println(settings7)
```

## group by

```go
	resul := make([]result, 0)
	
	db.Debug().Model(User{}).Select("Name, sum(Age) as Total").Group("Name").Find(&resul)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1\
	for k, v := range resul {
		fmt.Println(k, v)
	}

	resul := make([]result, 0)

	db.Debug().Model(User{}).Select("Name, sum(Age) as Total").Group("Name").Find(&User{}, "Name = ?", "小王子").Scan(&resul)
//等价于
	db.Debug().Model(User{}).Select("Name, sum(Age) as Total").Group("Name").Where("Name = ?", "小王子").Scan(&resul)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1\
	for k, v := range resul {
		fmt.Println(k, v)
	}


```



## Distinct

```go
var user []User
//name 和age 同时满足才能重复
db.Debug().Distinct("name", "age").Find(&user)
var user2 []User
//这个user2 只有name 被接收了
db.Debug().Distinct("name").Find(&user2)
log.Println(user)
log.Println(user2)
```

## Scan

Scan 结果至 struct，用法与 `Find` 类似

```go
var result []User
db.Table("users").Select("name", "age").Where("name = ?", "小王子").Scan(&result)
log.Println(result)

// Raw SQL
var result2 []User
db.Raw("SELECT name, age FROM users WHERE name = ?", "氧气").Scan(&result2)
log.Println(result2)
```

# 更新

## 更新单个列

```go
// 条件更新
db.Debug().Model(&User{}).Where("name = ?", "氧气").Update("age", 123)
//  UPDATE `users` SET `age`=123 WHERE name = '氧气'

// User 的 ID 是 `111`
var user2 User
user2.ID = 3
db.Debug().Model(&user2).Update("age", "20")
// UPDATE `users` SET `age`='20' WHERE `id` = 3

// 根据条件和 model 的值进行更新
var user User
user.ID = 6
db.Debug().Model(&user).Where("name = ?", "焦糖").Update("age", "1212")
//  UPDATE `users` SET `age`='1212' WHERE name = '焦糖' AND `id` = 6
```



## 更新多列

`Updates` 方法支持 `struct` 和 `map[string]interface{}` 参数。当使用 `struct` 更新时，默认情况下，GORM 只会更新非零值的字段

```go
// 根据 `struct` 更新属性，只会更新非零值的字段
var user User
user.ID = 4
db.Debug().Model(&user).Updates(User{Name: "Miya", Age: 18})
// UPDATE `users` SET `name`='Miya',`age`=18 WHERE `id` = 4

// 根据 `map` 更新属性
var user2 User
user2.ID = 5
db.Debug().Model(&user2).Updates(map[string]interface{}{"name": "氧气", "age": 18})
// UPDATE `users` SET `age`=18,`name`='氧气' WHERE `id` = 5
```



## Select和Omit

选中和忽略

```go
// 使用 Map 进行 Select
var user User
user.ID = 2
//只更新选中的name 字段
db.Debug().Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
//  UPDATE `users` SET `name`='hello' WHERE `id` = 2

var user2 User
user2.ID = 3
//忽略name 字段的
db.Debug().Model(&user2).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
//  UPDATE `users` SET `age`=18 WHERE `id` = 3
```





## 更新 Hook

GORM 支持的 hook 点包括：`BeforeSave`, `BeforeUpdate`, `AfterSave`, `AfterUpdate`. 更新记录时将调用这些方法，查看 [Hooks](https://gorm.io/zh_CN/docs/hooks.html) 获取详细信息

```go
    var user User
   user.ID = 2
   //只更新选中的name 字段
   db.Debug().Model(&user).Select("name").Updates(map[string]interface{}{"name": "氧气", "age": 18})
   //  UPDATE `users` SET `name`='hello' WHERE `id` = 2

}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
   if u.ID == 2 {
      log.Fatalln("admin user not allowed to update")
      return errors.New("admin user not allowed to update")
   }
   return
}
```



## 批量更新

```go
//批量更新
db.Debug().Model(User{}).Where("age=?", 1212).Updates(User{Name: "半糖app", Age: 19})
// [rows:3] UPDATE `users` SET `name`='半糖app',`age`=19 WHERE age=18


	// 根据 map 更新
db.Debug().Table("users").Where("id IN ?", []int{2, 3}).Updates(map[string]interface{}{"name": "氧气app", "age": 18})
	//  UPDATE `users` SET `age`=18,`name`='氧气app' WHERE id IN (2,3)
```





## 阻止全局更新

如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 `ErrMissingWhereClause` 错误

对此，你必须加一些条件，或者使用原生 SQL，或者启用 `AllowGlobalUpdate` 模式，例

```go
//报错
err2 := db.Debug().Model(&User{}).Update("name", "jinzhu").Error // gorm.ErrMissingWhereClaus  (WHERE conditions required )
fmt.Println(err2)
 
//成功
db.Debug().Model(&User{}).Where("1 = 1").Update("name", "羊羊")
//  [rows:5] UPDATE `users` SET `name`='羊羊' WHERE 1 = 1   //全表更新

	db.Debug().Exec("UPDATE users SET name = ?", "氧气")
	//  [rows:5] UPDATE users SET name = '氧气'  //全表更新成功

	db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("age", 100)
	//  [rows:5] UPDATE `users` SET `age`=100  //全表更新成功
```

## 更新的记录数

获取受更新影响的行数

就是用一个变量去接受最后返回的结果

```go
result := db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("age", 101)
fmt.Println(result.RowsAffected) //影响行数
fmt.Println(result.Error)
```

## 高级选项

### 使用 SQL 表达式更新

GORM 允许使用 SQL 表达式更新列，例如：

```go
user := &User{
   ID: 2,
}
	result := db.Debug().Model(&user).Update("age", gorm.Expr("age * ?+?", 2, 100))
	// UPDATE `users` SET `age`=age * 2+100 WHERE `id` = 2
	result2 := db.Debug().Model(&User{ID: 3}).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 100)})
	// UPDATE `users` SET `age`=age * 2 + 100 WHERE `id` = 3
	result3 := db.Debug().Model(&User{ID: 4}).UpdateColumn("age", gorm.Expr("age-?", 1))
	// UPDATE `users` SET `age`=age-1 WHERE `id` = 4

```

### 子查询

有问题



# 删除

## 删除一条记录

删除一条记录时，删除对象需要指定主键，否则会触发 [批量 Delete](https://gorm.io/zh_CN/docs/delete.html#batch_delete)，例如：

```go
	db.AutoMigrate(&User{})
	db.Create([]User{
		{Name: "氧气语音", Age: 12},
		{Name: "氧气语音2", Age: 12},
	})
	var user User
	user.ID = 1
	db.Debug().Delete(&user)
	//只能对Id 有效
	db.Debug().Where("name=?", "氧气语音2").Delete(&User{})
	// 顺序不能调返
```



```go
db.AutoMigrate(&Product{})
db.Create(&Product{ProductName: "产品", UserId: 3})
var product []Product
db.Find(&product)
log.Println(product)
```

## 根据主键删除

GORM 允许通过主键(可以是复合主键)和内联条件来删除对象，它可以使用数字（如以下例子。也可以使用字符串——译者注）。查看 [查询-内联条件（Query Inline Conditions）](https://gorm.io/zh_CN/docs/query.html#inline_conditions) 了解详情。

```go
db.Debug().Delete(&User{}, 5)
// DELETE FROM `users` WHERE `users`.`id` = 5
db.Debug().Delete(&User{}, []int{7, 8})
// DELETE FROM `users` WHERE `users`.`id` IN (7,8
```



## 批量删除

如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录

```go
db.Debug().Where("name like ?", "%氧气%").Delete(&User{})
//DELETE FROM `users` WHERE name like '%氧气%'
db.Debug().Delete(&User{}, "name like ?", "%氧%")
//DELETE FROM `users` WHERE name like '%氧%'
```



### 阻止全局删除

如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 `ErrMissingWhereClause `错误

对此，你必须加一些条件，或者使用原生 SQL，或者启用 `AllowGlobalUpdate` 模式，例如

```go
err2 := db.Delete(&User{}).Error
log.Fatalln(err2)
	db.Debug().Where("1=1").Delete(User{})
	//全表删除

	db.Debug().Exec("delete from users")
	//全表删除
	db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
	//全表删除
```

### 返回删除行的数据

返回被删除的数据，仅适用于支持 Returning 的数据库，例如：

```go
// 返回所有列
var users []User
DB.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING *
// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

// 返回指定的列
DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
```



## 软删除

如果您的模型包含了一个 `gorm.deletedat` 字段（`gorm.Model` 已经包含了该字段)，它将自动获得软删除的能力！

拥有软删除能力的模型调用 `Delete` 时，记录不会从数据库中被真正删除。但 GORM 会将 `DeletedAt` 置为当前时间， 并且你不能再通过普通的查询方法找到该记录。

```go
	db.Debug().Delete(&User{ID: 1})
	//UPDATE `users` SET `deleted_at`='2022-11-09 18:01:33.068' WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	db.Debug().Where("age", 12).Delete(&User{})
	//[rows:3] UPDATE `users` SET `deleted_at`='2022-11-09 18:03:40.375' WHERE `age` = 12 AND `users`.`deleted_at` IS NULL

	var userList []User
	db.Debug().Find(&userList)
	//SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL 自动忽略
	for _, user := range userList {
		fmt.Println(user)
	}
```

如果您不想引入 `gorm.Model`，您也可以这样启用软删除特性：

```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

### 查找被软删除的记录

您可以使用 `Unscoped` 找到被软删除的记录

```go
var userList2 []User
db.Debug().Unscoped().Where("age=12").Find(&userList2)
// SELECT * FROM `users` WHERE age=12
```

### 永久删除

您也可以使用 `Unscoped` 永久删除匹配的记录

```go
//永久删除
db.Unscoped().Delete(&User{ID: 1})

for _, user := range userList2 {
   fmt.Println(user)
}
```

# Belogs to

## Has one

:star: 一对一关系

```go
// User 有一张 CreditCard，UserID 是外键
type User struct {
   gorm.Model
   CreditCard CreditCard
}

type CreditCard struct {
   gorm.Model
   Number string
   UserID uint
}
```



```go
    users, err := GetAll(db)
   for i, user := range users {
      fmt.Println(i, user)
   }
}

func GetAll(db *gorm.DB) ([]User, error) {
   var user = []User{}
   err := db.Debug().Model(&User{}).Preload("CreditCard").Find(&user).Error
   //SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
   return user, err
}
```





### 重写外键

```go
type User struct {
   gorm.Model
   CreditCard CreditCard `gorm:"foreignKey:UserName"` // 使用 UserName 作为外键
}

type CreditCard struct {
   gorm.Model
   Number   string
   UserName string
}
```

### 多态关联

```go
type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}




db.AutoMigrate(&Toy{})
db.AutoMigrate(&Cat{})
db.AutoMigrate(&Dog{})
db.Debug().Create(&Cat{Toy: Toy{Name: "toy1"}, Name: "cat1"})
// 对于cat 表： INSERT INTO `cats` (`name`) VALUES ('cat1')
// 对于toys表：  INSERT INTO `toys` (`name`,`owner_id`,`owner_type`) VALUES ('toy1',1,'cats') ON DUPLICATE KEY UPDATE `owner_type`=VALUES(`owner_type`),`owner_id`=VALUES(`owner_id`)
```



查询

```go
func GetAll(db *gorm.DB) ([]Cat, error) {
   var cat = []Cat{}
   err := db.Debug().Model(&Cat{}).Preload("Toy").Find(&cat).Error
   // SELECT * FROM `toys` WHERE `owner_type` = 'cats' AND `toys`.`owner_id` = 1
   //SELECT * FROM `cats`
   return cat, err
}
```

## Has Many

一对多

```go
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

db.AutoMigrate(CreditCard{})
db.AutoMigrate(User{})
var users = &User{
   CreditCards: []CreditCard{
      {Number: "359"},
      {Number: "258"},
      {Number: "147"},
   },
}
//如果想的话，可以单独多的表增加需要的记录

db.Debug().Create(&users)
//对于 CreditCards 表 [rows:3]
// 对于users 表  [rows:1] 


//查询方法
func GetAll(db *gorm.DB) ([]User, error) {
	var users = []User{}
	err := db.Debug().Model(&User{}).Preload("CreditCards").Find(&users).Error
	// [rows:3] SELECT * FROM `credit_cards` WHERE `credit_cards`.`user_id` = 1 AND `credit_cards`.`deleted_at` IS NULL
	//[rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL
	return users, err
}

	usersList, err := GetAll(db)
	for i, user := range usersList {
		fmt.Println(i, user)
	}
```



## 多对多

```go
// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

db.Debug().Create(&User{
   Languages: []Language{
      {Name: "中文"},
      {Name: "英文"},
   },
})
// [rows:2]   INSERT INTO `languages` (
//[rows:2]   INSERT INTO `user_languages`
//[rows:1] INSERT INTO `users` 

//查询
func GetAll(db *gorm.DB) ([]User, error) {
	var users = []User{}
	err := db.Debug().Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
}
```

## 预加载

```go
package main

import (
   "gorm.io/driver/mysql"
   "gorm.io/gorm"
   "log"
)

// 文章
type Topics struct {
   Id         int        `gorm:"primary_key"`
   Title      string     `gorm:"not null"`
   UserId     int        `gorm:"not null"`
   CategoryId int        `gorm:"not null"`
   Category   Categories `gorm:"foreignkey:CategoryId"` //文章所属分类外键
   User       Users      `gorm:"foreignkey:UserId"`     //文章所属用户外键
}

// 用户
type Users struct {
   Id   int    `gorm:"primary_key"`
   Name string `gorm:"not null"`
}

// 分类
type Categories struct {
   Id   int    `gorm:"primary_key"`
   Name string `gorm:"not null"`
}

func main() {
   //db.InitMySQL("xyim:U8eKl90thLK2pMj@tcp(192.168.3.112:13306)", "xyim_sys", 10)

   dsn := "xyim:U8eKl90thLK2pMj@tcp(192.168.3.112:13306)/xyim_sys?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil {
      log.Fatalln(err)
   }
   //创建表
   //db.Debug().AutoMigrate(&Categories{}, &Users{}, &Topics{})

   //创建记录
   //var topis = &Topics{
   // Title:    "博客1",
   // Category: Categories{Name: "开发"},
   // User:     Users{Name: "作者：框加"},
   //}
   //db.Create(&topis)

   //预加载查询
   var topic Topics
   db.Debug().Where("id=?", 2).Preload("Category").Preload("User").First(&topic)
   //SELECT * FROM `categories` WHERE `categories`.`id` = 1
   // SELECT * FROM `users` WHERE `users`.`id` = 1
   //  SELECT * FROM `topics` WHERE id=2 ORDER BY `topics`.`id` LIMIT 1    博客ID=2
   log.Println(topic)
}
```



# 链式方法

被污染的：

```go
var userList = []User{}
var userList2 = []User{}
query := db.Where("name=?", "烤鸡")
query.Where("age>?", 10).Find(&userList)
query.Debug().Where("age>?", 0).Find(&userList2)
//SELECT * FROM `users` WHERE name='烤鸡' AND age>10 AND age>0
log.Println(userList)
log.Println(userList2)
```

解决：

```go
var userList = []User{}
var userList2 = []User{}
query := db.Debug().Where("name=?", "烤鸡").Session(&gorm.Session{})
query.Where("age>?", 10).Find(&userList)
//[rows:2] SELECT * FROM `users` WHERE name='烤鸡' AND age>10
query.Debug().Where("age>?", 0).Find(&userList2)
//[rows:3] SELECT * FROM `users` WHERE name='烤鸡' AND age>0
log.Println(userList)
log.Println(userList2)
```

#### NewDB

通过 `NewDB` 选项创建一个不带之前条件的新 DB，例如：

```go
var userList = []User{}
var userList2 = []User{}
query := db.Debug().Where("name=?", "烤鸡").Session(&gorm.Session{NewDB: true})
query.Debug().Where("age>?", 10).Find(&userList)
//SELECT * FROM `users` WHERE age>10
query.Debug().Where("age>?", 0).Find(&userList2)
//SELECT * FROM `users` WHERE age>0
log.Println(userList)
log.Println(userList2)
```







