/**
 * @Author Shershon
 * @Description 连接数据库
 * @Date 2021/7/8 6:08 下午
 **/
package gormpkg

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//// gorm.Config
//var gormConfig = &gorm.Config{
//	// 命名策略
//	NamingStrategy: schema.NamingStrategy{
//		// 表前缀
//		TablePrefix: "mc_",
//		// 使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
//		SingularTable: true,
//	},
//}

// 默认连接
func ConnectMysqlByDefault(host, port, user, pass, dbname string) (*gorm.DB, error) {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return gorm.Open(mysql.Open(dns), &gorm.Config{})
}

// 自定义配置连接
func ConnectMysqlByCustom(host, port, user, pass, dbname string) (*gorm.DB, error) {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //用change重命名列，MySQL8之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}))
}

// 基于database/sql连接，进行二次封装
func ConnectMysqlByDatabaseSql(host, port, user, pass, dbname string) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.New(mysql.Config{Conn: db}))
}
