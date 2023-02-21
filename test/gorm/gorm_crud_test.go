/**
 * @Description 测试使用gorm crud的使用
 **/
package tgorm

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"shershon1991/go-tools/app/gormpkg"
	"testing"
	"time"
)

var mysqlClient *gorm.DB

// 初始化mysql客户端
func init() {
	// 连接客户端
	host := "127.0.0.1"
	user, pass, port, database := "root", "root", "3306", "test"
	databaseSql, _ := gormpkg.ConnectMysqlByDefault(host, port, user, pass, database)
	mysqlClient = databaseSql
}

// 保存一条记录
func TestAddOne(t *testing.T) {
	// 初始化结构体
	userRow := gormpkg.User{
		NickName:     "李四",
		Age:          18,
		Phone:        "12340000",
		MemberNumber: "A0001",
		Birthday:     sql.NullString{String: "1991-03-04", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	}
	// 传入指针
	result := mysqlClient.Create(&userRow)
	fmt.Println(result)
	fmt.Println("id: ", userRow.ID)
	fmt.Println("插入记录错误: ", result.Error)
	fmt.Println("插入记录的条数: ", result.RowsAffected)
}

// 批量插入
func TestBatchInsert(t *testing.T) {
	// 定义user 切片
	userRows := []gormpkg.User{
		{NickName: "路人甲", Age: 20, Phone: "20000000", MemberNumber: "A0002"},
		{NickName: "路人乙", Age: 22, Phone: "30000000", MemberNumber: "A0003"},
		{NickName: "路人丙", Age: 24, Phone: "40000000", MemberNumber: "A0004"},
		{NickName: "路人丁", Age: 25, Phone: "50000000", MemberNumber: "A0005"},
	}
	// 保存
	result := mysqlClient.Create(&userRows)
	fmt.Println("插入记录错误: ", result.Error)
	fmt.Println("插入记录的条数: ", result.RowsAffected)
	// 打印ID
	for _, row := range userRows {
		fmt.Println("插入ID: ", row.ID)
	}
}

// 使用提供的First、Take、Last，查询单条记录
func TestGetOne(t *testing.T) {
	// 定义对应的结构体变量存储结果
	var firstUser gormpkg.User
	var takeUser gormpkg.User
	var lastUser gormpkg.User
	var result *gorm.DB
	// 获取第一条记录（主键升序） SELECT * FROM users ORDER BY id LIMIT 1;
	result = mysqlClient.First(&firstUser)
	fmt.Printf("Affected Rows: %+v, First Result: %+v\n", result.RowsAffected, firstUser)
	// 获取一条记录，没有指定排序字段 SELECT * FROM users LIMIT 1;
	result = mysqlClient.Take(&takeUser)
	fmt.Printf("Affected Rows: %+v, Take Result: %+v\n", result.RowsAffected, takeUser)
	// 获取最后一条记录（主键降序）SELECT * FROM users ORDER BY id DESC LIMIT 1;
	result = mysqlClient.Last(&lastUser)
	fmt.Printf("Affected Rows: %+v, LastUser Result: %+v\n", result.RowsAffected, lastUser)
}

// 使用Find(默认查询的是检索全部)
func TestGetByFind(t *testing.T) {
	var userList []gormpkg.User
	// 指定查询字段
	result := mysqlClient.Select("id", "nick_name").Find(&userList)
	//fmt.Printf("userList: %+v \n", userList)
	for _, user := range userList {
		fmt.Printf("id: %d nick_name: %s \n", user.ID, user.NickName)
	}
	fmt.Println("查询记录数: ", result.RowsAffected)
	fmt.Println("查询错误: ", result.Error)
}

// 根据String条件查询
func TestGetByStringWhere(t *testing.T) {
	// 定义对应的结构体变量
	var user gormpkg.User
	var userList []gormpkg.User
	var result *gorm.DB
	// 字符串条件查询一条
	result = mysqlClient.Where("nick_name = ?", "李四").First(&user)
	fmt.Printf("Res1: %v err:%v \n", user, result.Error)
	// 字符串条件查询多条
	result = mysqlClient.Where("nick_name <> ?", "李四").Find(&userList)
	fmt.Printf("Res2: %v err:%v \n", userList, result.Error)
	// 多个条件
	result = mysqlClient.Where("nick_name = ? and age >= ?", "李四", 18).First(&user)
	fmt.Printf("Res3: %v err:%v \n", user, result.Error)
}

// 字符串条件连表查询
func TestGetUserJoinAddress(t *testing.T) {
	// 定义对应的结构体变量
	var userAdd gormpkg.UserJoinAddress
	var userAddList []gormpkg.UserJoinAddress
	var result *gorm.DB
	// 连表查询一条
	result = mysqlClient.Table("users AS u").
		Joins("LEFT JOIN user_addresses AS ua ON u.id=ua.uid").
		Where("u.nick_name = ?", "老王").
		Select("u.id,u.created_at,u.updated_at,u.deleted_at,u.nick_name,ua.uid,ua.province,ua.city,ua.area,ua.detail").
		First(&userAdd)
	fmt.Printf("Res1: %+v err:%+v \n", userAdd, result.Error)
	// 连表查询多条
	result = mysqlClient.Table("users AS u").
		Joins("LEFT JOIN user_addresses AS ua ON u.id=ua.uid").
		Where("u.id >= ?", 2).
		Select("u.id,u.created_at,u.updated_at,u.deleted_at,u.nick_name,ua.uid,ua.province,ua.city,ua.area,ua.detail").
		Find(&userAddList)
	fmt.Printf("Res2: %+v err:%+v \n", userAddList, result.Error)
}

// 根据struct和map 条件查询结果
func TestGetByStructAndMapWhere(t *testing.T) {
	// 定义对应的结构体变量
	var user gormpkg.User
	var userList []gormpkg.User
	var result *gorm.DB
	// 结构体条件
	result = mysqlClient.Where(&gormpkg.User{NickName: "李四", Age: 18}).First(&user)
	fmt.Printf("结构体条件: %+v err:%v \n", user, result.Error)
	// map条件
	result = mysqlClient.Where(map[string]interface{}{"age": 18}).Find(&userList)
	fmt.Printf("map条件: %+v err:%v \n", userList, result.Error)
	// 主键切片
	result = mysqlClient.Where([]int64{2, 3, 4, 5}).Find(&userList)
	fmt.Printf("主键切片: %+v err:%v \n", userList, result.Error)
}

// 更新单个字段
func TestUpdateColumn(t *testing.T) {
	var result *gorm.DB
	// 字符串条件更新
	// UPDATE users SET nick_name='张三A', updated_at=当前时间 WHERE nick_name='张三;
	result = mysqlClient.Model(&gormpkg.User{}).Where("nick_name = ?", "李四").
		Update("nick_name", "张三A")
	fmt.Printf("条件更新: %+v err:%v \n", result.RowsAffected, result.Error)
	// 结构体条件更新
	// UPDATE users SET age=28, updated_at=当前时间 WHERE member_number='A0001;
	result = mysqlClient.Model(&gormpkg.User{}).Where(&gormpkg.User{MemberNumber: "A0001"}).Update("age", 28)
	fmt.Printf("结构体条件更新: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 更新多个字段
func TestUpdateMultipleColumn(t *testing.T) {
	var result *gorm.DB
	// 使用map
	updateMap := map[string]interface{}{
		"age":      32,
		"birthday": "1991-01-05",
	}
	// UPDATE users SET age=32,birthday='1991-01-05',updated_at=当前时间 WHERE id=6;
	result = mysqlClient.Model(&gormpkg.User{}).Where("id = ?", 4).Updates(updateMap)
	fmt.Printf("使用map结构更新: %+v err:%v \n", result.RowsAffected, result.Error)
	// 使用结构体(不使用Select)
	updateUser := gormpkg.User{
		Birthday: sql.NullString{String: "1993-10-10", Valid: true},
		Age:      0,
	}
	// @注意这里的age=0不会更新到MySQL
	// UPDATE users SET birthday='1993-09-09',updated_at=当前时间 WHERE id=5;
	result = mysqlClient.Model(&gormpkg.User{}).Where("id = ?", 5).Updates(updateUser)
	fmt.Printf("使用struct结构更新: %+v err:%v \n", result.RowsAffected, result.Error)
	// 使用结构体(使用Select)
	updateUser2 := gormpkg.User{
		Birthday: sql.NullString{String: "1993-09-09", Valid: true},
		Age:      0,
	}
	// UPDATE users SET birthday='1993-09-09',age=0,updated_at=当前时间 WHERE id=4;
	result = mysqlClient.Model(&gormpkg.User{}).
		Select("birthday", "age"). //指定要更新的字段
		Where("id = ?", 3).Updates(updateUser2)
	fmt.Printf("使用struct结构更新2: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 删除数据(软删除)
func TestSoftDel(t *testing.T) {
	var result *gorm.DB
	// 根据主键，删除一条记录
	result = mysqlClient.Delete(&gormpkg.User{}, 1)
	fmt.Printf("根据主键删除一条: %+v err:%v \n", result.RowsAffected, result.Error)
	// 根据主键切片，删除多条记录
	result = mysqlClient.Delete(&gormpkg.User{}, []int64{2, 3})
	fmt.Printf("根据主键切片删除多条: %+v err:%v \n", result.RowsAffected, result.Error)
	// 根据条件删除
	result = mysqlClient.Where("age = ?", 32).Delete(&gormpkg.User{})
	fmt.Printf("根据条件删除: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 删除数据(硬删除)
func TestStrongDel(t *testing.T) {
	var result *gorm.DB
	result = mysqlClient.Unscoped().Delete(&gormpkg.User{}, 1)
	fmt.Printf("硬删除: %+v err:%v \n", result.RowsAffected, result.Error)
}

// 事务使用
func TestTransaction(t *testing.T) {
	err := mysqlClient.Transaction(func(tx *gorm.DB) error {
		//在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 创建用户记录
		user := gormpkg.User{NickName: "老王", Age: 48}
		if err := tx.Create(&user).Error; err != nil {
			// 回滚事务
			return err
		}
		// 创建用户地址
		userAddress := gormpkg.UserAddress{Uid: user.ID, Province: "北京", City: "北京", Area: "海淀区"}
		if err := tx.Create(&userAddress).Error; err != nil {
			// 回滚事务
			return err
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("执行完成")
}

// 手动事务
func TestUseManualTx(t *testing.T) {
	// 用户表
	user := gormpkg.User{NickName: "小丽", Age: 19}
	// 开启事务
	tx := mysqlClient.Begin()
	// 添加用户
	if err := tx.Create(&user).Error; err != nil {
		// 遇到错误时回滚事务
		fmt.Println("添加用户失败: ", err)
		tx.Rollback()
	}
	// 用户地址表
	userAddress := gormpkg.UserAddress{Uid: user.ID, Province: "北京", City: "北京", Area: "昌平区"}
	// 添加用户地址
	if err := tx.Create(&userAddress).Error; err != nil {
		// 遇到错误时回滚事务
		fmt.Println("添加用户地址失败: ", err)
		tx.Rollback()
	}
	// 提交事务
	tx.Commit()
	fmt.Println("执行完成")
}
