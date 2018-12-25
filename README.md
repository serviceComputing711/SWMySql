# SWMySql
database in StarWars with MySql

## 1 数据库
### 1.1 数据库使用

> 数据库使用方法在`MySql.go`文件中有详细注释。不赘述

### 1.2 用户登录

> 创建数据库名称为`SW`，登录名为`testuser`，密码为`123456`。
> 直接使用即可，用户名和密码已在代码中，只需调用函数。
> 由于封装在`Docker`中，封装名称为`mysql`，需要先启动，才可以工作。

### 1.3 数据来源

> 修改上一版中的`database.getInfor.go`爬虫，从`SWAPI`中将数据记录下来。

### 1.4 功能测试

> 直接使用外层的`testDB.go`即可。
