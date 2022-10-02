/**
 * @Description 测试使用gorm迁移
 **/
package gormpkg

// 自动迁移schema,(根据结构体创建或者更新schema)
func GormAutoMigrate(host, port, use, pass, database string) error {
	mysqlByDefault, err := ConnectMysqlByDefault(host, port, use, pass, database)
	if err != nil {
		return err
	}
	// 指定引擎和表备注
	//err = mysqlByDefault.Set("gorm:table_options", "ENGINE=InnoDB COMMENT='用户地址表'").AutoMigrate(&UserAddress{})

	err = mysqlByDefault.Set("gorm:table_options", "ENGINE=InnoDB COMMENT='用户表'").Migrator().AutoMigrate(&UserList{})
	if err != nil {
		return err
	}
	return nil
}
